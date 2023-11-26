# go-mock-api

A simple mock api server written in go with rate limiting and authentication with redis as a database.

## Usage

### docker compose

run the server and db

```bash
make run # or docker compose up -d
```

stop the server and db

```bash
make stop # or docker compose down
```

Now you can access the server at http://localhost:3000
