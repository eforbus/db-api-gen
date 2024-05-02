#!/usr/bin/env bash
set -euo pipefail

echo "==> Running service."
# run service on localhost:3000
go run -mod=mod main.go
