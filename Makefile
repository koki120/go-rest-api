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
