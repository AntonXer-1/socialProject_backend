include .env
export

migrate_up:
	@migrate -path migrations -database ${CONN_STRING} up 1

migrate_down:
	@migrate -path migrations -database ${CONN_STRING} down 1