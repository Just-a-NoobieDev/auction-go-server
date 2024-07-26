.PHONY: run migrate dev db-up db-down db-init

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=Cjroxas14!
DB_NAME=auctiondb
MIGRATIONS_DIR=/docker-entrypoint-initdb.d

run:
	@echo "Running the application..."
	go run cmd/main.go

migrate:
	@echo "Applying database migrations..."
	docker-compose exec postgres psql -h ${DB_HOST} -U $(DB_USER) -d $(DB_NAME) -a -f $(MIGRATIONS_DIR)/initial.sql

dev:
	@echo "Starting development server with live reloading..."
	rm -rf ./tmp && air -c .air.toml
	

db-up:
	@echo "Starting PostgreSQL container..."
	docker-compose up -d

db-down:
	@echo "Stopping PostgreSQL container..."
	docker-compose down

db-init: db-up
	@echo "Waiting for PostgreSQL to start..."
	@sleep 10 # wait for PostgreSQL to initialize
	@echo "Applying initial migrations..."
	$(MAKE) migrate
