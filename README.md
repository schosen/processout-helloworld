This is a simple hello world HTTP server

## Pre-requisites:

You must have the following installed to run locally.
- Docker and docker-compose
- go
- minikube
- kubectl
- Terraform

To download dependencies run:
```
go mod tidy
```


<!-- To run this app
```
go run main.go
``` -->

## Docker

Docker-compose is configured for local development to run this application:
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

This application has prometheus monitoring. The following metrics are:

- **requestsTotal**: a counter that tracks the total number of HTTP requests received by the server
- **responseDuration**: a histogram that measures the duration of HTTP responses
-  **errorTotal**: a counter that tracks total number of errors returned by server (the errors have been simulated using a random boolean function. When set to true, the handler returns an "Internal Server Error" and increments the errorsTotal counter)

If I had more time I would add cpu / memory usage metrics
<!-- add more on metrics on container memory / cpu usage  -->

## CI/ CD
This application uses github action for its continous integration pipeline, the Docker image is pushed to docker hub when new code is merged to main

TODO: Add how we deploy to minikube in the pipeline


## Kubernetes deployment
This application can be deployed to minikube cluster.

Minikube was chosen as it sets up a single node Kubernetes cluster on your local machine which is a good alternative if you don't have acess to cloud services to test.


I've shared two ways to deploy;

Option 1 Manifests:
I have created Kubernetes Manifests allows you to specify the desired state of your application which can be deployed into your Kubernetes cluster. To deploy using manifests:

Start you minikube cluster
```
minikube start
```
jump into manifest directory
```
cd manifests
```

make executable
```
chmod +x deploy.sh
```

run the deploy.sh script
```
./deploy.sh
```

To see the deployed application run
```
kubectl get pods -n checkout
```

If it's running successfully you should see this
```
NAME                                     READY   STATUS    RESTARTS   AGE
processout-helloworld-5bbbbf9c8d-7jc5x   1/1     Running   0          99s
```


Option 2 terraform:
Terraform is used to create a application resources. I personally think this isnt needed / overkill for local kubernetes deployment the however if using cloud computing service like aws eks IaC would be used to provision resources. It's a tooling I don't have experience with so wanted to give it a try.

Start you minikube cluster
```
minikube start
```

To initialise terraform run
```
terraform init
```

To preview config before applying run
```
terraform plan
```

To apply config run
```
terraform apply
```

# Accessing the kubernetes pod locally
