.DEFAULT_GOAL := default

# Default target
default: tidy build

# Build the project
build:
	@echo "Building..."
	@if [ ! -d "./bin" ]; then mkdir bin; fi
	@go build -o bin/thumbtack cmd/thumbtack/main.go

# Install the project
install:
	@go install

# Update the project modules and tidy
update: updatemods tidy

# Tidy the project modules
tidy:
	@echo "Making mod tidy"
	@go mod tidy

# Update the project modules
updatemods:
	@echo "Updating..."
	@go get -u ./...

# Test the project
test:
	@echo "Testing..."
	@go test -covermode=atomic ./...