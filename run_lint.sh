#!/bin/bash
set -e

# Install golangci-lint if not present
if [ ! -f "./bin/golangci-lint" ]; then
    echo "Installing golangci-lint..."
    make golangci-lint
fi

# Run golangci-lint
echo "Running golangci-lint..."
./bin/golangci-lint run --timeout 10m