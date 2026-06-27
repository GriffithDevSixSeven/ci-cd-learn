include .env

export



MIGRATION_PATH=./migrations
DOCKER_NETWORK=ci_cd_net
MIGRATE_CMD=docker run --rm -v $(shell pwd)/$(MIGRATION_PATH):/migrations --network $(DOCKER_NETWORK) migrate/migrate -path=/migrations -database "$(MIGRATE_URL)"
.PHONY: migrate-create migrate-up migrate-down migrate-force



db-up:
	@docker compose up -d postgres-service --remove-orphans

db-down:
	@docker compose down postgres-service


migrate-create:
	@read -p "Введите имя миграции: " name; \
	$(MIGRATE_CMD) create -ext sql -dir /migrations -seq $$name

migrate-up:
	@$(MIGRATE_CMD) up

migrate-down:
	@$(MIGRATE_CMD) down 1

migrate-force:
	@read -p "Введите версию: " version ;\
	$(MIGRATE_CMD) force $$version