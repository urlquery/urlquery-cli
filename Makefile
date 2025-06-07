# Makefile for urlquery-cli

# Variables
BINARY_NAME=urlquery-cli
VERSION?=dev
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT}"

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	go build ${LDFLAGS} -o ${BINARY_NAME} .

# Build for multiple platforms
.PHONY: build-all
build-all:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BINARY_NAME}-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-windows-amd64.exe .

# Run tests
.PHONY: test
test:
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint code
.PHONY: lint
lint:
	golangci-lint run

# Clean build artifacts
.PHONY: clean
clean:
	rm -f ${BINARY_NAME}*
	rm -f coverage.out coverage.html

# Install the binary
.PHONY: install
install: build
	go install ${LDFLAGS} .

# Run the application
.PHONY: run
run:
	go run ${LDFLAGS} . $(ARGS)

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        - Build the binary"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install the binary"
	@echo "  run          - Run the application (use ARGS=... for arguments)"
	@echo "  help         - Show this help"