#!/bin/bash

# Navigate to workspace
cd /workspace

# Use gofmt to fix formatting
echo "Running gofmt on Go files..."
gofmt -w pkg/api/v1beta1/awspcaissuer_types.go
gofmt -w pkg/aws/pca.go

echo "Formatting complete."

# Check if files are properly formatted
echo "Checking formatting..."
gofmt -d pkg/api/v1beta1/awspcaissuer_types.go
gofmt -d pkg/aws/pca.go

echo "Done."