#!/bin/bash
set -e

echo "Converting spaces to tabs in awspcaissuer_types.go..."
# Convert leading spaces to tabs (assuming 8 spaces = 1 tab for Go)
sed -i 's/^        /\t/g' pkg/api/v1beta1/awspcaissuer_types.go

echo "Converting spaces to tabs in pca.go..."
# Convert leading spaces to tabs (assuming 8 spaces = 1 tab for Go)
sed -i 's/^        /\t/g' pkg/aws/pca.go

echo "Running gofmt to ensure proper formatting..."
gofmt -w pkg/api/v1beta1/awspcaissuer_types.go
gofmt -w pkg/aws/pca.go

echo "Checking formatting..."
echo "Checking awspcaissuer_types.go:"
gofmt -d pkg/api/v1beta1/awspcaissuer_types.go

echo "Checking pca.go:"
gofmt -d pkg/aws/pca.go

echo "Done."