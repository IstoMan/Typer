# Makefile for the typer project

BINARY_NAME=typer
BIN_DIRECTORY=bin

.PHONY: all build run clean

all: build

# Builds the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BIN_DIRECTORY)/$(BINARY_NAME) main.go

# Runs the binary
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BIN_DIRECTORY)/$(BINARY_NAME)

# Cleans the binary
clean:
	@echo "Cleaning..."
	@rm -f $(BIN_DIRECTORY)/$(BINARY_NAME)
