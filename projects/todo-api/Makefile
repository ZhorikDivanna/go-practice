include .env
export

DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable

postgres-up:
	docker compose up postgres

run-local:
	go run main.go

run-docker:
	docker compose up --build app

migrate-up:
	migrate -path migrations/ -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations/ -database "$(DATABASE_URL)" down 1