#!/bin/bash

IMAGE_NAME=${1:-snippetbox}
TAG=${2:-latest}
CLUSTER_NAME=$(kind get clusters) # Get the name of the kind cluster, assuming there is only one

# Load the image into the kind cluster

kind load docker-image $IMAGE_NAME:$TAG --name "$CLUSTER_NAME"
