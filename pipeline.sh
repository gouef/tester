#!/usr/bin/env bash
set -e

echo "Running tests..."
go test ./...

echo "Building application..."
go build -o myapp

echo "Build complete."
