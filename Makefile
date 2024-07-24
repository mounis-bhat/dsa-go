# Define the executable name
BINARY_NAME=dsa-go

# Default target to build the binary
all: build

# Build the binary
build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME)

# Run the binary
run: build
	@echo "Running the project..."
	./$(BINARY_NAME)

# Clean up the binary
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Lint the code
lint:
	@echo "Linting the code..."
	gofmt -l -s -w .

# Test the code
test:
	@echo "Running tests..."
	go test ./...

# Phony targets to avoid conflicts with files named clean, build, etc.
.PHONY: all build run clean deps lint test
