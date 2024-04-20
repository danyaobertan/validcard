# Valicard Project

## Introduction
The Valicard project is a robust web application designed to validate credit card information efficiently. Utilizing Docker, it offers a seamless deployment experience while adhering to clean code and architecture principles for maintainability and scalability.

## Features
- Validate credit card numbers using the Luhn algorithm.
- Verify expiration dates to ensure validity.
- Docker integration for simplified deployment and enhanced isolation.

## Requirements
- Docker
- Go (if running without Docker)

## Installation
Begin by cloning the repository to your local machine:
```bash
git clone https://github.com/danyaobertan/validcard.git

cd valicard
```

## Makefile
Take advantage of the Makefile to streamline operations:
```bash
make help
```
Output:
```bash
Usage:
  make install       - Install all dependencies.
  make run           - Run the application.
  make build         - Build the executable binary.
  make test          - Run tests.
  make lint          - Run linter.
  make clean         - Clean the binary.
  make docker-build  - Build the Docker image.
  make docker-run    - Run the Docker container.
  make docker-stop   - Stop and remove the Docker container.
  make docker-clean  - Remove Docker image.
  make docker-compose- Run the Docker container using docker-compose.
  make help          - Show this help message.
```

## Luhn Algorithm
The Luhn algorithm validates credit card numbers, implemented in `pkg/lunn.go`. It has been rigorously tested on 10,000 credit card numbers generated via https://www.validcreditcardnumber.com.

## CI/CD 
The project employs a robust CI/CD pipeline using GitHub Actions, triggered on every push to the main branch.

## Testing
Extensive testing ensure the code's correctness. Execute tests with:
```bash
make test
```

## golangci-lint
Utilize golangci-lint was used to enforce best practices:
```bash
make lint
``` 

## API Documentation
The project incorporates Swagger for API documentation. 
Alternatively, you can use some of requests in the `requests.http` file.