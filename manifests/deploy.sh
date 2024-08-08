#!/bin/bash

# Ensure script fails if any command fails
set -e

# applying the application to my cluster
kubectl apply -f namespace.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
