include .env

export

db-up:
	@docker compose up -d postgres-service --remove-orphans

db-down:
	@docker compose down postgres-service
