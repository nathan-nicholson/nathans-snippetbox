#!/bin/bash

# Fail fast
set -eou pipefail

./scripts/build.sh

./scripts/load-kind-with-image.sh

helm upgrade my-snippetbox charts/snippetbox -n my-snippetbox --install --values=./config/helm/values.yaml --create-namespace

kubectl rollout restart deployment -n my-snippetbox
