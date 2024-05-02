#!/usr/bin/env bash
set -euo pipefail

export DATABASE_WAIT_SECONDS="20"
export MYSQL_DATABASE="${MYSQL_DATABASE:-employees}"
export MYSQL_ROOT_USER="${MYSQL_ROOT_USER:-root}"
export MYSQL_ROOT_PASSWORD="${MYSQL_ROOT_PASSWORD:-root}"
export SCHEMA_FILE="${SCHEMA_FILE:-./schemas/employees.sql}"
export SCHEMA_DIR="$(dirname ${SCHEMA_FILE})"
export SERVICE_NAME="$(basename $(pwd))"

echo "==> Stopping database."
docker-compose down -v --remove-orphans
