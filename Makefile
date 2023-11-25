build:
	@go build -o bin/go-mock-api main.go

test:
	@go test -v ./...

run:
	@docker-compose up -d

stop:
	@docker-compose down

docker-build:
	@docker build -t go-mock-api .


build-run:
	@make docker-build
	@make run