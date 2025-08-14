# Makefile for the typer project

BINARY_NAME=typer

.PHONY: all build run clean

all: build

# Builds the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) main.go

# Runs the binary
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

# Cleans the binary
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
