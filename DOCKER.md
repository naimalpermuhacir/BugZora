# BugZora Docker Guide

This guide explains how to run BugZora using Docker containers, including all optimizations, security features, and advanced usage.

## 🚀 Docker Highlights

- **Multi-stage build**: Small, production-ready images
- **Alpine Linux base**: Minimal and secure
- **Trivy installation**: Latest release, direct from GitHub
- **Non-root user**: Container runs as UID 1000
- **Read-only root filesystem**: Enhanced security
- **Dropped capabilities**: Only essential Linux capabilities enabled
- **Health checks**: Dockerfile and Compose healthcheck support
- **Resource limits**: Memory and CPU limits in Compose
- **Volume caching**: Trivy cache for faster scans
- **Proper labels**: OCI and Docker metadata
- **Multi-arch build script**: `build-docker.sh` for amd64, arm64, arm/v7
- **Security scan script**: `docker-security-scan.sh` for automated vulnerability/config/secret scan

## Requirements

- Docker
- Docker Compose (recommended)

## Quick Start

### Using Docker Compose (Recommended)

```bash
git clone https://github.com/naimalpermuhacir/BugZora.git
cd BugZora
docker-compose build
docker-compose run --rm bugzora --help
```

### Using Docker directly

```bash
docker build -t bugzora:latest .
docker run --rm bugzora:latest --help
```

### Multi-Architecture Build

```bash
./build-docker.sh v1.2.0
```

### Security Scan of Container

```bash
./docker-security-scan.sh v1.2.0
ls -la security-scan-results/
```

## Features

- **Multi-stage build**: Smaller final image size
- **Non-root user**: Security-first approach
- **Trivy integration**: Automatic Trivy installation (latest release)
- **Volume mounting**: Host filesystem scanning capability
- **Docker socket access**: Docker image scanning capability
- **Lightweight base**: Alpine-based with minimal footprint
- **Security hardened**: Read-only filesystem, security options, dropped capabilities
- **Resource limits**: Memory and CPU limits in Compose
- **Health checks**: Dockerfile and Compose healthcheck support
- **Multi-arch support**: Build and run on amd64, arm64, arm/v7

## Security Features

- Container runs as non-root user (`bugzora`)
- Read-only volume mounting for filesystem scans
- Security options enabled (`no-new-privileges`, `seccomp`)
- Minimal base image reduces attack surface
- Only essential Linux capabilities enabled
- Health checks for container liveness

## Usage Examples

### Different Output Formats

```bash
# JSON output
docker-compose run --rm bugzora image alpine:latest --output json

# PDF output
docker-compose run --rm bugzora image ubuntu:20.04 --output pdf

# Table output (default)
docker-compose run --rm bugzora image debian:11 --output table

# Quiet mode
docker-compose run --rm bugzora image nginx:alpine --quiet
```

### Filesystem Scanning

```bash
# Scan current directory
docker-compose run --rm bugzora fs /scan-target

# Scan specific path
docker run --rm -v /path/to/scan:/target:ro bugzora:latest fs /target

# Scan with different output formats
docker-compose run --rm bugzora fs /scan-target --output json
docker-compose run --rm bugzora fs /scan-target --output pdf
```

### Private Registry Scanning

```bash
# Login to private registry
docker login registry.example.com

# Scan private image
docker-compose run --rm bugzora image registry.example.com/myapp:v1.0.0
```

### Advanced Scanning Options

```bash
# Scan with specific severity levels
docker-compose run --rm bugzora image nginx:latest --severity HIGH,CRITICAL

# Skip unfixed vulnerabilities
docker-compose run --rm bugzora image alpine:latest --ignore-unfixed

# Include all packages
docker-compose run --rm bugzora image ubuntu:20.04 --list-all-pkgs

# Offline scanning
docker-compose run --rm bugzora fs /scan-target --offline-scan
```

## Docker Compose Configuration

The `docker-compose.yml` file includes:

- **Resource limits**: Memory and CPU constraints
- **Volume caching**: Trivy cache for faster scans
- **Security options**: Read-only filesystem, security options
- **Health checks**: Container health monitoring
- **Environment variables**: Configurable settings

```yaml
version: '3.8'
services:
  bugzora:
    build: .
    image: bugzora:latest
    container_name: bugzora-scanner
    user: "1000:1000"
    volumes:
      - ./:/scan-target:ro
      - trivy-cache:/tmp/.trivy-cache
    working_dir: /scan-target
    environment:
      - TRIVY_CACHE_DIR=/tmp/.trivy-cache
    security_opt:
      - no-new-privileges:true
    cap_drop:
      - ALL
    healthcheck:
      test: ["CMD", "bugzora", "--help"]
      interval: 30s
      timeout: 10s
      retries: 3
    deploy:
      resources:
        limits:
          memory: 2G
          cpus: '2.0'
        reservations:
          memory: 512M
          cpus: '0.5'

volumes:
  trivy-cache:
```

## Dockerfile Features

The `Dockerfile` includes:

- **Multi-stage build**: Separate build and runtime stages
- **Alpine Linux**: Minimal base image
- **Trivy installation**: Latest release from GitHub
- **Non-root user**: Security-first approach
- **Security hardening**: Read-only filesystem, dropped capabilities
- **Health checks**: Container health monitoring
- **Proper labels**: OCI and Docker metadata

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bugzora .

# Runtime stage
FROM alpine:3.18
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bugzora .

# Install Trivy
RUN wget -qO - https://github.com/aquasecurity/trivy/releases/download/v0.50.1/trivy_0.50.1_Linux-64bit.tar.gz | tar -xz -C /usr/local/bin trivy

# Create non-root user
RUN addgroup -g 1000 bugzora && \
    adduser -D -s /bin/sh -u 1000 -G bugzora bugzora

USER bugzora
ENTRYPOINT ["./bugzora"]
```

## Multi-Architecture Support

Build and run on different architectures:

```bash
# Build for multiple architectures
./build-docker.sh v1.2.0

# Available architectures:
# - linux/amd64 (x86_64)
# - linux/arm64 (ARM64)
# - linux/arm/v7 (ARMv7)
```

## Security Scanning

Automated security scanning of the container:

```bash
# Run security scan
./docker-security-scan.sh v1.2.0

# View results
ls -la security-scan-results/
cat security-scan-results/summary_*.md
```

## Troubleshooting

### Common Issues

1. **Permission denied**
   ```bash
   # Fix volume permissions
   chmod 755 /path/to/scan
   ```

2. **Trivy not found**
   ```bash
   # Rebuild image
   docker-compose build --no-cache
   ```

3. **Memory issues**
   ```bash
   # Increase memory limit
   docker-compose run --rm -m 4g bugzora image alpine:latest
   ```

### Debug Mode

```bash
# Enable debug output
docker-compose run --rm bugzora image alpine:latest --debug

# Enable trace logging
docker-compose run --rm bugzora fs /scan-target --trace
```

## Performance Optimization

- **Volume caching**: Trivy cache for faster scans
- **Resource limits**: Memory and CPU constraints
- **Multi-stage build**: Smaller final image size
- **Alpine base**: Minimal footprint

## Best Practices

1. **Use read-only volumes** for filesystem scans
2. **Set resource limits** to prevent resource exhaustion
3. **Use non-root user** for security
4. **Enable health checks** for monitoring
5. **Cache Trivy database** for faster scans
6. **Use multi-arch builds** for different platforms
