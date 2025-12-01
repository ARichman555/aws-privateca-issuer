#!/bin/bash
set -e

echo "Running gofmt on all Go files..."
gofmt -w .

echo "Done."