DB_DSN := "postgres://postgres:postgres@localhost:5432/warrior?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)


migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	${MIGRATE} down

run:
	@docker compose up --build -d
	@go run cmd/main.go