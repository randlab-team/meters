# Meters service
Web api and importing cli tool for meters logs.

## Building

### Binaries
Build binaries to the `build/` directory:

```
$ make build
```

### Docker

```
$ docker build -f docker/prod/meters_service.dockerfile -t <TAG> .
```
Build test image 
```
$ make docker-dev-build
```

## Running

### Binaries

#### Web api

```
$ export DB_STRING=postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable
$ export ALLOWED_ORIGINS=*
$ ./build/meters_service
```

#### Import tool

```
$ ./build/meters_import -d <DB_STRING> csv -f <PATH_TO_CSV>
```
Example (needs db):

``` 
$ make test-import
```

### Docker compose

Run whole dev env with fresh docker build:
```
$ make dev-env-up
```

Run db:
```
$ make dev-env-up-db
```

## Web api documentation

### Endpoints

#### Get all meter logs
Request:
```
GET http://{{host}}/v1/meters/
```
Response 200 - found one or more meters:
```json
[
  {
    "id": int,
    "sn": int,
    "correct": bool,
    "param_name": string,
    "index": int,
    "date_register": date,
    "value": int,
    "log_interval": int,
    "status": int
  },
  {...}
]
```
Response 204 - no meters found

Response 500 - error during handling request

## Database

### Schema
`sql/meters_table.sql`

```SQL
CREATE TABLE meters
(
    id            SERIAL PRIMARY KEY,
    sn            INTEGER,
    correct       BOOL,
    param_name    VARCHAR(16),
    "index"       INTEGER,
    date_register TIMESTAMP,
    "value"       INTEGER,
    log_interval  INTEGER,
    status        INTEGER
);

CREATE UNIQUE INDEX meters_unique_sn_date_reg
    ON meters (sn, date_register);
```