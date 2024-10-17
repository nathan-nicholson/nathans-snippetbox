#!/bin/bash

# Create a new kind cluster
kind create cluster --config=config/kind_config.yaml

# Install the Nginx Ingress Controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for the Nginx Ingress Controller to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s


# Build and deploy the snippetbox application for initial load
sh scripts/deploy.sh
