#!/bin/bash

# Get the cluster name from kind_config

CLUSTER_NAME=$(yq '.name' ./config/kind_config.yaml)

kind delete cluster --name "$CLUSTER_NAME"
