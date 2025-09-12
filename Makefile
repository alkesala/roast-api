.PHONY: build run clean test dev install

# Variables
BINARY_NAME=roast-api
GO_FILES=$(shell find . -name "*.go" -type f)

# Build the application
build:
	go build -o $(BINARY_NAME) .

# Run the application
run: build
	./$(BINARY_NAME)

# Run in development mode (auto-reload with air if available)
dev:
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "Air not found. Install with: go install github.com/cosmtrek/air@latest"; \
		go run .; \
	fi

# Install dependencies
install:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test -v ./...

# Run with race detection
test-race:
	go test -race -v ./...

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install from https://golangci-lint.run/usage/install/"; \
	fi

# Vet code
vet:
	go vet ./...

# Check code quality
check: fmt vet lint test

# Docker commands
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 --env-file .env $(BINARY_NAME)

# Fly.io deployment
deploy:
	flyctl deploy

# Show help
help:
	@echo "Available targets:"
	@echo "  build        Build the application"
	@echo "  run          Build and run the application"
	@echo "  dev          Run in development mode"
	@echo "  install      Install dependencies"
	@echo "  clean        Clean build artifacts"
	@echo "  test         Run tests"
	@echo "  test-race    Run tests with race detection"
	@echo "  fmt          Format code"
	@echo "  lint         Lint code"
	@echo "  vet          Vet code"
	@echo "  check        Run all code quality checks"
	@echo "  docker-build Build Docker image"
	@echo "  docker-run   Run Docker container"
	@echo "  deploy       Deploy to Fly.io"
	@echo "  help         Show this help"