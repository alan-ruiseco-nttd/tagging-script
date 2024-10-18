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

.PHONY: tools-install
tools-install:
	asdf pluging add golangci-lint || true
	asdf pluging add pre-commit || true
	asdf install

.PHONY: config
config: tools-install
	@echo '#!/bin/sh' > .git/hooks/pre-commit
	@echo 'golangci-lint run' >> .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "Pre-commit hook has been setup to run golangci-lint"