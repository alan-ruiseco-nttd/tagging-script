# Makefile

# Define the binary name
BINARY_NAME=main_tagging

# Run the Go application
.PHONY: run
run: build
	./$(BINARY_NAME)

# Clean up build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)