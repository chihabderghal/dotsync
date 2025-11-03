# Get the current version from git tag or commit hash
VERSION := $(shell git describe --tags --always --abbrev=0)

# Binary names
BINARY := ./bin/dotsync
BINARY_VERSIONED := ./bin/dotsync-$(VERSION)

# Build binary with version embedded
build:
	@echo "Building $(BINARY_VERSIONED)..."
	go build -o $(BINARY_VERSIONED) -ldflags "-X main.version=$(VERSION)" ./cmd/dotsync/main.go
	@echo "Built $(BINARY_VERSIONED)"

# Build default binary without version suffix
build-default:
	@echo "Building $(BINARY)..."
	go build -o $(BINARY) -ldflags "-X main.version=$(VERSION)" ./cmd/dotsync/main.go
	@echo "Built $(BINARY)"

# Run CLI with arguments
run:
	go run ./cmd/dotsync/main.go $(ARGS)

# Clean binaries
clean:
	rm -f ./bin/dotsync ./bin/dotsync-*

.PHONY: build build-default run clean
