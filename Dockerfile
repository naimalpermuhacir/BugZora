# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with security flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -installsuffix cgo \
    -ldflags="-w -s -extldflags '-static'" \
    -o bugzora .

# Runtime stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    wget \
    curl \
    && rm -rf /var/cache/apk/*

# Install Trivy
RUN wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | \
    apk add --no-cache --repository https://aquasecurity.github.io/trivy-repo/deb/ trivy

# Create non-root user
RUN addgroup -g 1000 bugzora && \
    adduser -D -s /bin/sh -u 1000 -G bugzora bugzora

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bugzora .

# Create necessary directories
RUN mkdir -p /tmp/trivy-cache /scan-target && \
    chown -R bugzora:bugzora /app /tmp/trivy-cache /scan-target

# Switch to non-root user
USER bugzora

# Set environment variables
ENV TRIVY_CACHE_DIR=/tmp/trivy-cache \
    TZ=UTC \
    BUGZORA_VERSION=1.2.0

# Add health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD ./bugzora --help > /dev/null 2>&1 || exit 1

# Add labels
LABEL maintainer="BugZora Team <bugzora@bugzora.dev>" \
      version="1.2.0" \
      description="Security scanning tool for container images and filesystems" \
      org.opencontainers.image.title="BugZora" \
      org.opencontainers.image.description="Comprehensive security scanning tool" \
      org.opencontainers.image.version="1.2.0" \
      org.opencontainers.image.vendor="BugZora" \
      org.opencontainers.image.source="https://github.com/naimalpermuhacir/BugZora"

# Expose port (for future API server)
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["./bugzora"]

# Default command
CMD ["--help"] 