include ../.env
DOCKER_COMMAND=docker compose -f ../docker-compose.yml
MIGRATE=${DOCKER_COMMAND} exec blog migrate -path=migrations -database "postgres://${BlogDBUsername}:${BlogDBPassword}@${BlogDBHost}:${BlogDBPort}/${BlogDBName}?sslmode=disable" -verbose

migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create-migration:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -seq -dir migrations  $$NAME
