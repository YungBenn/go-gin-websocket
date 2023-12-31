statik:
	statik -src=./docs/swagger -dest=./docs

docker:
	docker compose up -d

run:
	go run ./cmd/api/main.go

docker.run:
	docker run -p 8000:8000 -p 9000:9000 rubenadi/workshop-web-golang:v1

docker.build:
	docker build -t rubenadi/workshop-web-golang:v1 .