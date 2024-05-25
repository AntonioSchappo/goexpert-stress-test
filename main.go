package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func worker(url string, jobs <-chan int, results chan<- int) {
	for range jobs {
		resp, err := http.Get(url)
		if err != nil {
			//number of calls with errors will be sinalyzed as status 0
			results <- 0
			continue
		}
		results <- resp.StatusCode
		resp.Body.Close()
	}
}

func main() {
	url := flag.String("url", "", "URL to test")
	requests := flag.Int("requests", 1, "Number of requests to perform")
	concurrency := flag.Int("concurrency", 1, "Number of multiple requests to make at a time")

	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required")
		return
	}

	jobs := make(chan int, *requests)
	results := make(chan int, *requests)

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			worker(*url, jobs, results)
			wg.Done()
		}()
	}

	for i := 0; i < *requests; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	statusCodes := make(map[int]int)

	for result := range results {
		statusCodes[result]++
	}

	fmt.Printf("Total time: %v\n", time.Since(start).Round(time.Second))
	fmt.Printf("Requests: %d\n", *requests)

	for status, count := range statusCodes {
		fmt.Printf("HTTP %d: %d\n", status, count)
	}
}
