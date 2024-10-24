include .env
export $(shell sed 's/=.*//' .env)

POSTGRESQL_URL = ${DB_CONN_URL}

build:
	@mkdir -p ./dist
	@go build -o ./dist/ ./cmd/api/main.go

server:
	@go run ./cmd/api/main.go

docker:
	@docker compose up

migrate-up:
	@migrate -database ${POSTGRESQL_URL} -path ./migrations up

migrate-down:
	@migrate -database ${POSTGRESQL_URL} -path ./migrations down

.PHONY: server exp