# /bin/bash
include configs/.env

migrate_cmd := migrate -path migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DBNAME}?sslmode=${DB_SSLMODE}" -verbose

migrate:
	$(migrate_cmd) up

migrate_down:
	$(migrate_cmd) down

swagger-init:
	swag init --parseDependency -g docs/docs.go