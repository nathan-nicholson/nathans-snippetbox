#!/bin/bash

# Fail fast
set -eou pipefail

./scripts/build.sh

# Fail if the current context does not include kind (no oopsies with prod clusters)
kubectl config current-context | grep kind || (echo "Current context is not a kind cluster" && exit 1)

# Load the image into the kind cluster
./scripts/load-kind-with-image.sh

# Deploy the application using the helm chart
helm upgrade my-snippetbox charts/snippetbox -n my-snippetbox --install --values=./config/helm/values.yaml --create-namespace

# Restart the deployment to force the newly loaded image to be used
kubectl rollout restart deployment -n my-snippetbox

# Wait for the deployment to be ready
kubectl wait --namespace my-snippetbox \
  --for=condition=available deployment/my-snippetbox \
  --timeout=120s
