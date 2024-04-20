.PHONY: build test clean docker-build docker-run docker-stop docker-clean install help

# Variables
APP_NAME=valicard
DOCKER_TAG=$(APP_NAME):latest

# Install all dependencies
install:
	@echo "Installing dependencies..."
	go mod download

# Build the binary for the application
build: install
	@echo "Building $(APP_NAME)..."
	go build -o ./bin/$(APP_NAME) ./cmd/app

# Run tests 
test: install
	@echo "Running tests..."
	go test ./... -v

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run	

# Clean up any built binaries
clean:
	@echo "Cleaning up..."
	rm -rf ./bin

# Build the Docker image
docker-build: build
	@echo "Building Docker image..."
	docker build -t $(DOCKER_TAG) .

# Run the Docker container
docker-run:
	@echo "Running Docker container..."
	docker run --name $(APP_NAME) -p 8080:8080 -d $(DOCKER_TAG)

# Stop and remove the Docker container
docker-stop:
	@echo "Stopping Docker container..."
	docker stop $(APP_NAME)
	docker rm $(APP_NAME)

# Remove Docker image
docker-clean:
	@echo "Removing Docker image..."
	docker rmi $(DOCKER_TAG)

# Target to help with auto-documenting Makefile. Use 'make help' to see usage of targets.
help:
	@echo "Usage:"
	@echo "  make install       - Install all dependencies."
	@echo "  make build         - Build the executable binary."
	@echo "  make test          - Run tests."
	@echo "  make lint          - Run linter."
	@echo "  make clean         - Clean the binary."
	@echo "  make docker-build  - Build the Docker image."
	@echo "  make docker-run    - Run the Docker container."
	@echo "  make docker-stop   - Stop and remove the Docker container."
	@echo "  make docker-clean  - Remove Docker image."
	@echo "  make help          - Show this help message."
