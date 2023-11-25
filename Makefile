run :
	@go run main.go

build:
	@go build -o bin/go-mock-api main.go

test:
	@go test -v ./...