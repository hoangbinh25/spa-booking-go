.PHONY: up down logs migrate-version db-shell test

up:
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f

migrate-version:
	docker compose run --rm migrate version

db-shell:
	docker compose exec postgres psql -U postgres -d spa_booking

test:
	go test ./...
