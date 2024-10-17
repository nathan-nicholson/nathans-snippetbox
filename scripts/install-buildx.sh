#!/bin/bash

# Install docker buildx for use with docker/colima

# arm64 For M1 Macs, amd64 for Intel Macs
ARCH=$("./scripts/get-arch.sh")
VERSION=v0.17.1 # Check https://github.com/docker/buildx/releases for the latest version
curl -LO https://github.com/docker/buildx/releases/download/${VERSION}/buildx-${VERSION}.darwin-${ARCH}
mkdir -p ~/.docker/cli-plugins
mv buildx-${VERSION}.darwin-${ARCH} ~/.docker/cli-plugins/docker-buildx
chmod +x ~/.docker/cli-plugins/docker-buildx
docker buildx version # verify installation
