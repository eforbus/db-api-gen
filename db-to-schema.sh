#!/usr/bin/env bash
set -euo pipefail

export MYSQL_DATABASE="${MYSQL_DATABASE:-employees}"
export MYSQL_ROOT_USER="${MYSQL_ROOT_USER:-root}"
export MYSQL_ROOT_PASSWORD="${MYSQL_ROOT_PASSWORD:-root}"

# https://entgo.io/blog/2021/10/11/generating-ent-schemas-from-existing-sql-databases/
echo "==> Generating ent schema from database."
go run -mod=mod ariga.io/entimport/cmd/entimport -dsn "mysql://${MYSQL_ROOT_USER}:${MYSQL_ROOT_PASSWORD}@tcp(127.0.0.1:3306)/${MYSQL_DATABASE}"

echo "==> Generating api spec and service."
# generate openapi spec and service from schema
go generate ./...
