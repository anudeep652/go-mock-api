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

stop the server and db

```bash
docker compose down
```
