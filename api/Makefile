.PHONY: all build run test clean

all: build

build:
	@go build -o ./bin/main .

run: build
	@go run ./bin/main

test:
	@go test ./...

clean:
	@rm -rf ./bin
