.PHONY: all test lint run dev build
all: test build 

test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

dev:
	go run cmd/main.go

build:
	go mod tidy
	go mod vendor
	go build -o bin/gj cmd/main.go