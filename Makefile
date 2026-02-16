.PHONY: build run test clean deps

BIN_PATH := bin/hexlet-path-size

deps:
	go mod download

build: deps
	go build -o $(BIN_PATH) cmd/hexlet-path-size/main.go

run: build
	./$(BIN_PATH)

test: build
	@OUTPUT=$$(./$(BIN_PATH) --help | head -n 1); \
	if echo "$$OUTPUT" | grep -q "hexlet-path-size"; then \
		echo "✅ Test passed: help output is correct"; \
	else \
		echo "❌ Test failed: unexpected output"; \
		exit 1; \
	fi

clean:
	rm -rf bin/