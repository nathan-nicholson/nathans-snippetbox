#!/bin/bash

# Build the snippetbox image for the specified platform
# Usage: ./scripts/build.sh [platform]
# Example: ./scripts/build.sh linux/arm64

# Set the platform to the first argument or default to linux/arm64
PLATFORM=${1:-linux/arm64}

docker buildx build --build-arg BUILDPLATFORM="$PLATFORM" . -t snippetbox:latest
