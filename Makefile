include .env
export



GOOSE_CMD=goose
GOOSE_DIR=./db/migration

migrate-up:
	$(GOOSE_CMD) -dir $(GOOSE_DIR) postgres $(DB_URL) up

migrate-down:
	$(GOOSE_CMD) -dir $(GOOSE_DIR) postgres $(DB_URL) down

start:
	@go run main.go