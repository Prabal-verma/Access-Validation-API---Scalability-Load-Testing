FROM golang:1.21-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service

# Use a minimal alpine image
FROM alpine:3.18

WORKDIR /

# Copy the binary from builder
COPY --from=0 /auth-service /auth-service

# Expose port
EXPOSE 8080

# Run the service
CMD ["/auth-service"] 