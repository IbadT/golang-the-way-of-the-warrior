DB_DSN := "postgres://postgres:postgres@localhost:5432/warrior?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}


migrate:
	${MIGRATE} up

migrate-down:
	$(MIGRATE) down

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go


lint:
	golangci-lint run --color=always

run:
	go run cmd/main.go