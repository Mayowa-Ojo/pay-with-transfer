.PHONY: run gogen new-migration setup install-migrate

run: # run the server
	ENV=dev go run -race -v main.go

MIGRATIONS_DIR=migrations
MIGRATIONS_CONFIG=config.yml

new-migration: # create a new migration file
	if [ ! -d "$(MIGRATIONS_DIR)" ]; then mkdir $(MIGRATIONS_DIR); fi
	sql-migrate new --config $(MIGRATIONS_DIR)/$(MIGRATIONS_CONFIG) $(name)

migrate-up: # migrate database to most recent version
	set -a && source .env && set +a && \
	sql-migrate up --config $(MIGRATIONS_DIR)/$(MIGRATIONS_CONFIG)

migrate-dry: # migrate database to most recent version
	set -a && source .env && set +a && \
	sql-migrate up -dryrun --config $(MIGRATIONS_DIR)/$(MIGRATIONS_CONFIG)

migrate-down: # undo a single migration
	set -a && source .env && set +a && \
	sql-migrate down --config $(MIGRATIONS_DIR)/$(MIGRATIONS_CONFIG)

setup: install-migrate # install dependencies

install-migrate: # install sql-migrate @link: https://github.com/rubenv/sql-migrate#installation
ifeq (, $(shell which sql-migrate))
	echo "found missing binary, installing..."
	go install github.com/rubenv/sql-migrate/...@latest
endif

gogen: # generate files
	go generate -x -v ./...
