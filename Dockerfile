# Network Scanner Dockerfile

# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")" -o /scanner ./cmd/scanner

# Final stage
FROM alpine:latest

# Add certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Copy binary from build stage
COPY --from=builder /scanner /usr/local/bin/scanner

# Run as non-root user
RUN adduser -D scanner
USER scanner

ENTRYPOINT ["scanner"]
CMD ["--help"]
