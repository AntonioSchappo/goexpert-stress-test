# goexpert-stress-test

This project is a simple implementation of a stress tester using go.

This project is a cli tool that accepts 3 parameters:
- `--url`: The URL to test.
- `--requests`: The number of requests to send.
- `--concurrency`: The number of concurrent requests to send.

This project is based on concepts of concurrency and parallelism.

The tester makes use of channels and wait groups to ensure that all requests are completed before exiting.

Please check below the commands to run the project:

1 - Create a Docker build image of the program:
```sh
docker build -t stresstest .
```
2 - Run the Docker container based on the image created:
```sh
docker run stresstest --url=http://google.com --requests=1000 --concurrency=10
```
3 - As an alternate, it is also possible to run the program locally with the following command:
```sh
go run main.go --url=https://google.com --requests=100 --concurrency=10
```