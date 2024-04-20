# Builder stage: This stage installs build tools and dependencies.
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o valicard ./cmd/app

# Final stage: This stage builds the final image with the compiled Go binary.
FROM alpine:3.14

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/valicard .

# Copy the configuration directory into the container
COPY ./config /root/config

# Expose port 8080 to the outside world
EXPOSE 8080

# Environment variable to specify the configuration path
ENV CONFIG_PATH="/root/config"

# Command to run the executable
CMD ["./valicard"]
