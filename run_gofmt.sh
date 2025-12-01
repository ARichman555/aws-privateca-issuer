#!/bin/bash
set -e

echo "Checking if gofmt is available..."
which gofmt

echo "Running gofmt on awspcaissuer_types.go..."
gofmt -w pkg/api/v1beta1/awspcaissuer_types.go

echo "Running gofmt on pca.go..."
gofmt -w pkg/aws/pca.go

echo "Checking formatting..."
echo "Checking awspcaissuer_types.go:"
gofmt -d pkg/api/v1beta1/awspcaissuer_types.go

echo "Checking pca.go:"
gofmt -d pkg/aws/pca.go

echo "Done."