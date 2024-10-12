#!/bin/bash

# Build the snippetbox image for the specified platform
# Usage: ./scripts/build.sh [platform]
# Example: ./scripts/build.sh linux/arm64

# Set the platform to the first argument or default to linux/arm64
PLATFORM=$("./scripts/get-arch.sh")

docker buildx build --build-arg BUILDPLATFORM="$PLATFORM" . -t snippetbox:latest
