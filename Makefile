.PHONY: build run clean test

# Build the program
build:
	go build -o gotype main.go

# Run the program
run:
	go run main.go

# Clean build artifacts
clean:
	rm -f gotype

# Install dependencies
deps:
	go mod tidy

# Run tests (if any)
test:
	go test ./...

# Build and run
all: build run

# Help
help:
	@echo "Available targets:"
	@echo "  build  - Build the executable"
	@echo "  run    - Run the program directly"
	@echo "  clean  - Remove build artifacts"
	@echo "  deps   - Install/update dependencies"
	@echo "  test   - Run tests"
	@echo "  all    - Build and run"
	@echo "  help   - Show this help"
