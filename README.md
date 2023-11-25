# go-mock-api

A simple mock api server written in go with rate limiting and authentication with redis as a database.

## Usage

### docker compose

build the image

```bash
make docker-build
```

run the server and db

```bash
docker compose up -d
```

or just run the below make command to build the image and run the server and db

```bash
make build-run
```

stop the server and db

```bash
docker compose down #or just run `make stop`
```
