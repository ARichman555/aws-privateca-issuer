#!/bin/bash
cd /workspace

# Format the Go files using gofmt
echo "Formatting pkg/api/v1beta1/awspcaissuer_types.go..."
gofmt -w pkg/api/v1beta1/awspcaissuer_types.go

echo "Formatting pkg/aws/pca.go..."
gofmt -w pkg/aws/pca.go

echo "Done formatting files."