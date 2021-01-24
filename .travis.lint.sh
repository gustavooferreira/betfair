#!/bin/bash

set -e

# Examine code
go vet ./...

if [ -n "$(gofmt -l .)" ]; then
  echo "Go code is not formatted:"
  gofmt -d .
  exit 1
fi
