#!/bin/bash
set -e

echo "Running make generate..."
make generate

echo "Running make fmt..."
make fmt

echo "Running make vet..."
make vet

echo "Running make manifests..."
make manifests

echo "All commands completed successfully."