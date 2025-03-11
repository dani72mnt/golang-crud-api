# Use the official Golang image as the base image
FROM golang:1.24.0-alpine3.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY ../go.mod ../go.sum /app/

# Download the dependencies
RUN go mod download

RUN go mod tidy

# Copy the rest of the application code
COPY ../ ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o shop ./cmd/main.go

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/shop .
COPY docker/.env .env
# Expose the port your application runs on (8090)
EXPOSE 8090

# Command to run the application
CMD ["./shop"]