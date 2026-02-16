.PHONY: build run test clean deps

BIN_PATH := bin/hexlet-path-size

deps:
	go mod download

build: deps
	go build -o $(BIN_PATH) cmd/hexlet-path-size/main.go

run: build
	./$(BIN_PATH)

test: build
	@./$(BIN_PATH) --help | grep -q "hexlet-path-size" && \
	./$(BIN_PATH) --help | grep -q "print size of a file or directory" && \
	echo "✅ Test passed: help output is correct" || \
	(echo "❌ Test failed: unexpected output" && exit 1)

clean:
	rm -rf bin/