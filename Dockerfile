# Stage 1: Build the Go application
FROM golang:1.23.4 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Disable CGO and build a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Stage 2: Create a lightweight image for production
FROM alpine:latest

WORKDIR /app

# Install SSL certificates (for secure API calls, if needed)
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 7000

# Use ENTRYPOINT instead of CMD for better container behavior
ENTRYPOINT ["./main"]
