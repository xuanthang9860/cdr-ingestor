db-up:
	docker compose -f cmd/db/docker-compose.yaml up -d

db-down:
	docker compose -f cmd/db/docker-compose.yaml down

db-logs:
	docker compose -f cmd/db/docker-compose.yaml logs -f
go:
	go mod tidy

run:
	go run cmd/server/main.go