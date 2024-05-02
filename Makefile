# Copy this file to your project and customize it.

.DEFAULT_GOAL := help

ACTUAL_PWD := $(PWD)
DATE := $(shell date +%FT%T%z)
USER := $(shell whoami)

PROGRAM := db-api-gen
LICENSE := apache-2.0
PACKAGE := github.com/eforbus/$(PROGRAM)
URL     := https://$(PACKAGE)

.PHONY: all
all:  # Run generally applicable Makefile targets
all: clean build test

.PHONY: build
build:  # Build all artifacts from database schema
	@echo "==> Generating schema from database."
	./db-to-schema.sh

.PHONY: clean
clean:  # Clean temporary files.
	@echo "==> Cleaning temporary files."

.PHONY: test
test:  # Run tests
	@echo "==> Running tests."

.PHONY: database-up
database-up:  # Start and bootstrap database
	@echo "==> Starting database."
	./db-start.sh

.PHONY: database-down
database-down:  # Stop database
	@echo "==> Stopping database."
	./db-stop.sh

.PHONY: run
run:  # Running service
	@echo "==> Running."
	./run.sh

help:  # Print list of Makefile targets
	@for f in $(MAKEFILE_LIST); do grep -E ':  #' $${f} | grep -v 'LIST\|BEGIN' | \
	sort -u | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'; done
