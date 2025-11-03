.PHONY: help test test-verbose test-short bench bench-verbose race coverage clean lint lint-install fmt vet check-all

help:
	@echo "Available targets:"
	@echo "  test          - Run all tests"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  test-short    - Run tests in short mode"
	@echo "  bench         - Run all benchmarks"
	@echo "  bench-verbose - Run benchmarks with verbose output"
	@echo "  race          - Run tests with race detector"
	@echo "  coverage      - Run tests with coverage report"
	@echo "  lint          - Run golangci-lint"
	@echo "  lint-install  - Install golangci-lint"
	@echo "  fmt           - Format code with gofmt"
	@echo "  vet           - Run go vet"
	@echo "  clean         - Clean test cache and coverage files"
	@echo "  check-all     - Run all quality checks (fmt, vet, lint, test, race)"

test:
	@echo "Running tests..."
	go test ./homework/... -v

test-verbose:
	@echo "Running tests with verbose output..."
	go test ./homework/... -v -count=1

test-short:
	@echo "Running tests in short mode..."
	go test ./homework/... -short

bench:
	@echo "Running benchmarks..."
	go test ./homework/... -bench=. -benchmem

bench-verbose:
	@echo "Running benchmarks with verbose output..."
	go test ./homework/... -bench=. -benchmem -v

race:
	@echo "Running tests with race detector..."
	go test ./homework/... -race

coverage:
	@echo "Running tests with coverage..."
	go test ./homework/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean:
	@echo "Cleaning test cache and coverage files..."
	go clean -testcache
	rm -f coverage.out coverage.html

# Linting targets
lint:
	@echo "Running golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not found. Run 'make lint-install' to install it."; \
		exit 1; \
	fi

lint-install:
	@echo "Installing golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint is already installed"; \
	else \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2; \
	fi

fmt:
	@echo "Formatting code..."
	@files=$$(gofmt -s -l ./homework); \
	if [ -z "$$files" ]; then \
		echo "✅ All Go files are already formatted."; \
	else \
		echo "🛠  Formatting files..."; \
		echo "$$files" | xargs gofmt -s -w; \
		echo "Done."; \
	fi

vet:
	@echo "Running go vet..."
	go vet ./...

check-all:
	@echo "Running all quality checks..."
	@echo "\n=== Formatting code ==="
	$(MAKE) fmt
	@echo "\n=== Running go vet ==="
	$(MAKE) vet
	@echo "\n=== Running linter ==="
	$(MAKE) lint
	@echo "\n=== Running tests ==="
	$(MAKE) test
	@echo "\n=== Running race detector ==="
	$(MAKE) race
	@echo "\n=== All checks completed successfully! ==="