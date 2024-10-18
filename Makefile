include .env
export $(shell sed 's/=.*//' .env)

POSTGRESQL_URL = ${DB_CONN_URL}

server:
	@go run ./cmd/api/main.go

db:
	@docker compose up

migrate-up:
	@migrate -database ${POSTGRESQL_URL} -path ./migrations up

migrate-down:
	@migrate -database ${POSTGRESQL_URL} -path ./migrations down

.PHONY: server exp