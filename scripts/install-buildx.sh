#!/bin/bash

# Install docker buildx for use with docker/colima

ARCH=arm64 # change to 'amd64' for x86_64
VERSION=v0.17.1 # Check https://github.com/docker/buildx/releases for the latest version
curl -LO https://github.com/docker/buildx/releases/download/${VERSION}/buildx-${VERSION}.darwin-${ARCH}
mkdir -p ~/.docker/cli-plugins
mv buildx-${VERSION}.darwin-${ARCH} ~/.docker/cli-plugins/docker-buildx
chmod +x ~/.docker/cli-plugins/docker-buildx
docker buildx version # verify installation
