.PHONY: build run test clean

BIN_PATH := bin/hexlet-path-size

build:
	go build -o $(BIN_PATH) cmd/hexlet-path-size/main.go

run: build
	./$(BIN_PATH)

test: build
	@OUTPUT=$$(./$(BIN_PATH)); \
	if [ "$$OUTPUT" = "Hello from Hexlet!" ]; then \
		echo "✅ Test passed: output is correct"; \
	else \
		echo "❌ Test failed: unexpected output"; \
		exit 1; \
	fi

clean:
	rm -rf bin/