#!/usr/bin/env bash
set -euo pipefail

export DATABASE_WAIT_SECONDS="20"
export MYSQL_DATABASE="${MYSQL_DATABASE:-employees}"
export MYSQL_ROOT_USER="${MYSQL_ROOT_USER:-root}"
export MYSQL_ROOT_PASSWORD="${MYSQL_ROOT_PASSWORD:-root}"
export SCHEMA_FILE="${SCHEMA_FILE:-./schemas/employees.sql}"
export SCHEMA_DIR="$(dirname ${SCHEMA_FILE})"
export SERVICE_NAME="$(basename $(pwd))"

echo "==> Starting a clean docker-compose environment."
docker-compose pull
docker-compose down -v --remove-orphans
docker-compose up -d

echo "==> Waiting ${DATABASE_WAIT_SECONDS} seconds for database(s) to come online."
sleep ${DATABASE_WAIT_SECONDS}

echo "==> Loading the database schema(s)."
docker exec -w /opt -i -e SCHEMA_FILE=$(basename "${SCHEMA_FILE}") ${SERVICE_NAME}-mysql-1 \
  sh -c 'exec mysql -uroot -p"${MYSQL_ROOT_PASSWORD}" < "${SCHEMA_FILE}"'

#docker-compose down -v --remove-orphans
#exit 0
