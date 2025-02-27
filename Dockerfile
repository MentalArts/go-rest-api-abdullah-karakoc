# Base image
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Set CGO_ENABLED to 0 for cross-compilation
ENV CGO_ENABLED=0

# Build the Go application
RUN go build -o main .

# Start a new stage for the final image
FROM alpine:latest

# Copy the binary from the builder stage
WORKDIR /app
COPY --from=builder /app/main .

# Command to run the binary
CMD ["./main"]

# Expose the port the app runs on
EXPOSE 8080
