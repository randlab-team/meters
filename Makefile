.PHONY: build

build-meters-service:
	go build -o build/meters_service ./cmd/meters_service/

build-meters-import:
	go build -o build/meters_import ./cmd/meters_import/

build: build-meters-service build-meters-import

test:
	go test ./...

test-import: build
	./build/meters_import -d postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable csv -f docker/dev/data.csv

dev-env-up:
	docker compose -f docker/dev/docker-compose.yaml up --build

dev-env-down:
	docker compose -f docker/dev/docker-compose.yaml down

docker-dev-build:
	docker build -f docker/prod/meters_service.dockerfile -t meters_service:dev .