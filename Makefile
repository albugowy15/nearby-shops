migrate_up:
	@go run cmd/migrations/up/main.go

migrate_down:
	@go run cmd/migrations/down/main.go

dev:
	@air

run:
	@go run cmd/web/main.go

build:
	@go build -o ./tmp/main cmd/web/main.go

start:
	@./tmp/main

doc:
	@swag init -d "./" -g "cmd/web/main.go"

seeder:
	@go run cmd/seeder/main.go
