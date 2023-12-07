.PHONY: run build test

run:
	go run main.go

build:
	go build -o myapp main.go

test:
	go test ./...

test-coverage:
	go test -cover ./...
