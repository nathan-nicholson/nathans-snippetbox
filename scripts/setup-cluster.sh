#!/bin/bash

# Create a new kind cluster
kind create cluster --config=config/kind_config.yaml


# Build and deploy the snippetbox application for initial load
sh scripts/deploy.sh
