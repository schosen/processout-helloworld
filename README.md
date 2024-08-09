This is a simple hello world HTTP server

## Pre-requisites:

You must have the following installed to run locally.
- Docker and docker-compose
- Go
- Minikube
- Kubectl
- Terraform

To download dependencies run:
```
go mod tidy
```

## Docker

Docker-compose is configured for local development. This useful mainly for multi-container environments, so may be needed in future. To run this application:
```
docker-compose build

docker-compose up
```

Alternatively to build and run the docker container using the docker command
```
docker build -t processout-helloworld .

docker run -p 8080:8080 processout-helloworld
```

## API

Base URL `http://localhost:8080`


#### Get Hello World Message
- Endpoint: /
- Method: GET
- Description: Returns a simple "Hello, World!" message.
- Example Request:
    ```bash
    curl -X GET http://localhost:8080/api/
    ```
- Response
    - Status Code: `200 OK`
    - Body

    ```bash
    Hello, World!
    ```
#### Get Users
- Endpoint: /users
- Method: GET
- Description: Returns a list of users in JSON format.
- Example Request:
    ```bash
    curl -X GET http://localhost:8080/api/users
    ```
- Response:
    - Status Code: `200 OK`
    - Body:
    ```json
    [
        {
            "id": 1,
            "name": "Sarah Chosen",
            "job": "DevOps Engineer"
        },
        {
            "id": 2,
            "name": "John Doe",
            "job": "Builder"
        }
    ]
    ```

## Metrics
Prometheus metrics are exposed for this app. To view follow the above build and run commands and go to
`
http://localhost:8080/metrics
`

This application has prometheus monitoring. The following custom metrics are:

- **requestsTotal**: a counter that tracks the total number of HTTP requests received by the server
- **responseDuration**: a histogram that measures the duration of HTTP responses
-  **errorTotal**: a counter that tracks total number of errors returned by server (When not nil, the handler returns an "Internal Server Error" and increments the errorsTotal counter).

There are also "out the box" metrics related to memory consumption, cpu consumption, etc. e.g. `process_cpu_seconds_total` (Total user and system CPU time spent in seconds) and `process_virtual_memory_bytes` (Virtual memory size in bytes)

The above metrics are key to understanding the health of the service. Typically at minumum you want to measure the 4 golden signals for any service. Latency, Traffic, Errors and Saturation

## Tests
I have included test that checks whether the service returns status 200 code and checks whether the response body includes "Hello World"

To run tests:
```
go test -v
```

## CI/ CD
This application uses github action for its continous integration pipeline. When new code is merged to main this triggers the pipeline. Tests are run and then the Docker image is pushed to docker hub.

I then create a minikube cluster, the cluster is checked by running a kubectl commands. The pipeline then deploys the go server to the cluster using the manifests. This could be used if you want to run integration tests for the service later on. Ideally I would deploy to a cloud managed kubernetes service like AWS EKS which I currently don't have access to.

## Kubernetes deployment
This application can be deployed to minikube cluster.

Minikube was chosen as it sets up a single node Kubernetes cluster on your local machine which is a good alternative if you don't have access to cloud services to test.


I've shared two ways to deploy:

### Option 1 Manifests:
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


### Option 2 terraform:
Terraform is used to create application resources. I personally think this isnt needed / overkill for local kubernetes deployment and for such a small application however if using cloud computing service like aws eks, IaC could be used to provision resources such as the cluster itself. It's a technology I don't have experience with so wanted to give it a try.

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
terraform apply -var="docker_username=schosen"
```

# Accessing the kubernetes pod locally
To access the pod as if it were on your local machine you can port forward

Identify the pods name
```
# kubectl get pods -n checkout

NAME                                     READY   STATUS    RESTARTS   AGE
processout-helloworld-5bbbbf9c8d-7jc5x   1/1     Running   0          5m
```

Forward a local port (e.g., 8080) to the port on the pod (e.g., 8080):
```
kubectl port-forward pod/processout-helloworld-5bbbbf9c8d-7jc5x 8080:8080 -n checkout

```
# With more time...
- I would have deployed grafana to visualize the metrics
- I would improve test coverage. Apply unit testing to terraform to avoid user errors, this could then be applied as an integration step in the pipeline and the bash script replaced with terraform.
- Update API to include standard api and version prefix
