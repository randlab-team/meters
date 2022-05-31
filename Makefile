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

dev-env-up-db:
	docker compose -f docker/dev/docker-compose.yaml up postgres

dev-env-up:
	docker compose -f docker/dev/docker-compose.yaml up --build

dev-env-down:
	docker compose -f docker/dev/docker-compose.yaml down

docker-dev-build:
	docker build -f docker/prod/meters_service.dockerfile -t meters_service:dev .

docker-dev-run:
	# run db firs. see  e.g. dev-env-up-db
	docker run -p 8080:8080 -e DB_STRING=postgresql://postgres:postgres@postgres:5432/postgres?sslmode=disable -e ALLOWED_ORIGINS=* --network=dev_default -t meters_service:dev