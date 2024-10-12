#!/bin/bash

# Get the architecture for the current platform

ARCH=$(uname -m)

# Output linux/amd64 for x86_64 and linux/arm64 for aarch64

case $ARCH in
  x86_64)
    echo "linux/amd64"
    ;;
  aarch64)
    echo "linux/arm64"
    ;;
  *)
    echo "Unknown architecture: $ARCH" >&2
    exit 1
    ;;
esac
