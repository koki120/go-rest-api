ENV_GITLABCI_FILE := .env.gitlabci
ENV_GITLABCI = $(shell cat $(ENV_GITLABCI_FILE))

# App Server
.PHONY: run-dev
run-dev:
	docker compose -f docker-compose.dev.yml up --build

.PHONY: destroy
destroy:
	docker compose -f docker-compose.dev.yml -f docker-compose.test.yml down --volumes --remove-orphans
	rm -rf mysql/mysql-data

# MySQL
.PHONY: init-db
init-db:
	docker compose -f docker-compose.dev.yml run --rm api go run cmd/initdb/initdb.go
