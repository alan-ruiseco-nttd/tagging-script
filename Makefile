# Makefile

# Define the binary name
BINARY_NAME=main_tagging

# Install asdf if not already installed
.PHONY: install-asdf
install-asdf:
	if ! which asdf > /dev/null; then \
		git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.8.1; \
		echo '. $$HOME/.asdf/asdf.sh' >> ~/.bashrc; \
		echo '. $$HOME/.asdf/completions/asdf.bash' >> ~/.bashrc; \
		source ~/.bashrc; \
	fi

# Install Go using asdf based on .tool-versions
.PHONY: install-go
install-go: install-go
	if ! asdf plugin-list | grep -q 'golang'; then \
		asdf plugin-add golang https://github.com/kennyp/asdf-golang.git; \
	fi
	asdf install golang $$(awk '/^go / {print $$2}' .tool-versions)
	asdf global golang $$(awk '/^go / {print $$2}' .tool-versions)

# Configure the environment
.PHONY: configure
configure: install-asdf install-go	

# Build the Go application
.PHONY: build
build: install-go
	go build -o $(BINARY_NAME) main_tagging.go types.go

# Run the Go application
.PHONY: run
run: build
	./$(BINARY_NAME)

# Clean up build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)