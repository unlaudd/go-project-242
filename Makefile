.PHONY: build run test clean deps

BIN_PATH := bin/hexlet-path-size

deps:
	go mod download

build: deps
	go build -o $(BIN_PATH) cmd/hexlet-path-size/main.go

run: build
	./$(BIN_PATH)

test:
	go test -v ./...

clean:
	rm -rf bin/