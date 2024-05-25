FROM golang:1.22.2 as builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stresstest

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/stresstest .
ENTRYPOINT ["./stresstest"]