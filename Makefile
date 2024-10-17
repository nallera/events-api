# Variables
BINARY_NAME=events-api
SRC_DIR=./cmd/app
BUILD_DIR=./bin
GO=go

# Default target
.PHONY: all
all: tidy build

# Go mod tidy target
.PHONY: tidy
tidy:
	@echo "Running go mod tidy..."
	$(GO) mod tidy

# Build target
.PHONY: build
build:
	@echo "Building the binary..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Test target
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./...

# Run target
.PHONY: run
run: build
	@echo "Running the application..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Clean target
.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

# Help target
.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo "Targets:"
	@echo "  all      - Run go mod tidy and build the binary (default)"
	@echo "  tidy     - Run go mod tidy"
	@echo "  build    - Build the binary"
	@echo "  run      - Build and run the application"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean the build"
	@echo "  help     - Show this help message"