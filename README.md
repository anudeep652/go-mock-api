# go-mock-api

A simple mock api server written in go with rate limiting and authentication with redis as a database.

## Usage

### docker

build the image

```bash
docker build -t go-mock-api .
```

run the container

```bash
docker run -p 3000:3000 go-mock-api
```
