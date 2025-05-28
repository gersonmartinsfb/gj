.PHONY: all test lint run dev build
all: test lint build 

test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

dev:
	go run cmd/main.go

lint:
	golangci-lint run --config .golangci.yml
	@echo "Linting completed."

build:
	go mod tidy
	go mod vendor
	go build -o bin/cli cmd/main.go