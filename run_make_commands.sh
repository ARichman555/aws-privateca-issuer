#!/bin/bash

cd /workspace

echo "Running make generate..."
make generate

echo "Running make fmt..."
make fmt

echo "Running make vet..."
make vet

echo "Done."