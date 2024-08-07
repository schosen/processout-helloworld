This is a simple hello world HTTP server

Pre-requisites:

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

Docker

A docker-compose is configured for local development to run this application:
```
docker-compose build

docker-compose up
```

Alternivaely to build and run the docker container using the c=docker command
```
docker build -t processout-helloworld .

docker run -p 8080:8080 processout-helloworld
```

Metrics
This application we're monitoring two metrics via prometheus

- requestsTotal: a counter that tracks the total number of HTTP requests received by the server
- responseDuration: a histogram that measures the duration of HTTP responses

<!-- add more on metrics on container memory / cpu usage  -->
