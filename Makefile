DB_DSN := "postgres://postgres:postgres@localhost:5432/warrior?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)


migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	${MIGRATE} down

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
# `oapi-codegen` - обращение к нашей утилите oapi
# `-config openapi/.openapi`  - за конфиг берем файл `.openapi` из папки `openapi`
# `-include-tags tasks` - Генерируем описанные ручки под тегом tasks из файла `openapi.yaml`
# -package tasks openapi/openapi.yaml &gt; ./internal/web/tasks/api.gen.go - Генерирум все описанное по пути internal/web/tasks . 
# Этот путь вам нужно создать (Создать в папке internal папку web и в ней папку tasks)

# go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest


run:
	@docker compose up --build -d
	@go run cmd/main.go