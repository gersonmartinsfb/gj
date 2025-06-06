.PHONY: all test lint run dev build
all: test build 

test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

dev:
	go run cmd/main.go

clean:
	rm -rf bin
	rm -f coverage.out
	rm -f coverage.html

build: clean
	@mkdir -p bin
	go mod tidy
	go mod vendor
	go build -o bin/gj cmd/main.go