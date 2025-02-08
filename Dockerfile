# Stage 1: Build the Go application
FROM golang:1.23.5 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Stage 2: Create a lightweight image for production
FROM golang:1.23.5

WORKDIR /app

# Copy the compiled binary from builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 7000

# Command to run the application
CMD ["./main"]
