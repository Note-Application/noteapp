# Stage 1: Build the Go application
FROM golang:1.23.4 AS builder

WORKDIR /app

# Copy go modules and download dependencies (leveraging Docker cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Ensure the binary is built without CGO (for static linking)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Install any runtime dependencies (if needed)
RUN apk --no-cache add ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 7000

# Run the application
CMD ["./main"]
