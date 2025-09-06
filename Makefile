build:
	sudo docker build -t cdr-ingest .
tag:
	sudo docker tag cdr-ingest xuanthang9860/softswitch-cdr-ingestor:v1.0.0

db-up:
	docker compose -f cmd/db/docker-compose.yaml up -d

db-down:
	docker compose -f cmd/db/docker-compose.yaml down

db-logs:
	docker compose -f cmd/db/docker-compose.yaml logs -f

up:
	docker compose -f cmd/server/docker-compose.yaml up -d
down:
	docker compose -f cmd/server/docker-compose.yaml down
logs:
	docker compose -f cmd/server/docker-compose.yaml logs -f

reset:
	docker compose -f cmd/server/docker-compose.yaml down
	docker compose -f cmd/server/docker-compose.yaml up -d

go:
	go mod tidy

run:
	go run cmd/server/main.go