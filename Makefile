# Makefile

# Define the binary name
BINARY_NAME=main_tagging

# Install dependencies
.PHONY: deps
deps:
	go mod tidy

# Build the Go application
.PHONY: build
build: deps
	go build -o $(BINARY_NAME) main_tagging.go types.go

# Run the Go application
.PHONY: run
run: build
	./$(BINARY_NAME)

# Clean up build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)