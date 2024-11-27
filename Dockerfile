# Build stage
FROM golang:alpine AS builder

# Install git and build tools
RUN apk update && \
    apk add --no-cache git build-base

# Setup build directory
WORKDIR /build

# Copy only necessary files for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .
COPY .env .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy only the binary and env file from builder
COPY --from=builder /build/main .
COPY --from=builder /build/.env .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
