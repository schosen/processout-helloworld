This is a simple hello world HTTP server

## Pre-requisites:

You must have the following installed to run locally.
- Docker and docker-compose
- go

To download dependencies run:
```
go mod tidy
```


<!-- To run this app
```
go run main.go
``` -->

## Docker

A docker-compose is configured for local development to run this application:
```
docker-compose build

docker-compose up
```

Alternivaely to build and run the docker container using the docker command
```
docker build -t processout-helloworld .

docker run -p 8080:8080 processout-helloworld
```

## Metrics
Prometheus metrics are exposed for this app. To view follow the above build and run commands and go to
`
http://localhost:8080/metrics
`

This application has monitoring the following metrics via prometheus:

- **requestsTotal**: a counter that tracks the total number of HTTP requests received by the server
- **responseDuration**: a histogram that measures the duration of HTTP responses
-  **errorTotal**: a counter that tracks total number of errors returned by server (the errors have been simulated by random using a random boolean function. When set to true, the handler returns an "Internal Server Error" and increments the errorsTotal counter)

If I had more time I would add cpu / memory usage metrics
<!-- add more on metrics on container memory / cpu usage  -->

## CI/ CD

