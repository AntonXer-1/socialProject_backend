include .env
export

export PROJECT_ROOT=$(shell pwd)

docker-start:
	@docker compose up -d

docker-stop:
	@docker compose down

migrate-create:
	@if [-z "$(seq)"]; then \
  		echo "Отсутствует параметр seq" \
  		exit 1; \
  	fi; \
  		docker compose run --rm migrate \
  		create \
  		-ext sql \
  		-dir /migrations \
  		-seq "$(seq)"

migrate-pattern:
	@docker compose run --rm migrate \
	-path /migrations \
	-database ${CONN_STRING} \
	"$(action)"

migrate-up:
	@make migrate-pattern action=up

migrate-down:
	@make migrate-pattern action=down