# App Server
.PHONY: run
run:
	docker compose -f docker-compose.yml up --build

.PHONY: destroy
destroy:
	docker compose -f docker-compose.yml down --volumes --remove-orphans
	rm -rf mysql/mysql-data

# MySQL
.PHONY: init-db
init-db:
	docker compose -f docker-compose.yml run --rm api go run cmd/initdb/initdb.go

# Tools
.PHONY: tools
tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1 # go 1.19に対応するバージョン

# Lint, Format
.PHONY: lint
lint: tools
	golangci-lint run ./... --timeout=5m

.PHONY: format
format: tools
	golangci-lint run ./... --fix
