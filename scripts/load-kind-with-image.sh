#!/bin/bash

IMAGE_NAME=$1
TAG=${2:-latest}

# Load the image into the kind cluster

kind load docker-image $IMAGE_NAME:$TAG --name nathan-test
