format:
	go fmt ./...

lint:
	golangci-lint run --fix

run:
	go run cmd/main.go

all: format lint run