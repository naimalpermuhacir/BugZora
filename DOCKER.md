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
# Scan from private registry
docker-compose run --rm bugzora image registry.example.com/myapp:v1.0.0

# With authentication (if needed)
docker run --rm -v ~/.docker:/root/.docker:ro bugzora:latest image private.registry.com/app:latest
```

### Advanced Usage

```bash
# Interactive shell for debugging
docker-compose run --rm bugzora sh

# Check Trivy version
docker-compose run --rm bugzora sh -c "trivy --version"

# Update Trivy database
docker-compose run --rm bugzora sh -c "trivy image --download-db-only"
```

### Multi-Architecture Build

```bash
./build-docker.sh v1.2.0
```

### Security Scanning

```bash
./docker-security-scan.sh v1.2.0
ls -la security-scan-results/
```

## Docker Compose Configuration

See `docker-compose.yml` for all options. Highlights:

- **Security options**: `no-new-privileges`, `seccomp`, dropped capabilities
- **Resource limits**: Memory and CPU
- **Health check**: Container liveness
- **Read-only root filesystem**
- **Named volumes**: Trivy cache for performance
- **User**: UID 1000
- **Environment variables**: TZ, TRIVY_CACHE_DIR, etc.

## Environment Variables

You can customize the container behavior with environment variables:

```bash
# Set Trivy cache directory
docker-compose run -e TRIVY_CACHE_DIR=/custom/cache bugzora image ubuntu:20.04

# Set custom Trivy options
docker-compose run -e TRIVY_ARGS="--severity HIGH,CRITICAL" bugzora image alpine:latest
```

## Troubleshooting

### Docker daemon error
```bash
sudo systemctl start docker
docker info
```

### Permission errors
```bash
sudo usermod -aG docker $USER
newgrp docker
```

### Trivy cache issues
```bash
docker-compose run --rm bugzora sh -c "rm -rf /tmp/trivy-cache/*"
docker-compose down
docker-compose build --no-cache
```

### Volume mounting issues
```bash
ls -la /path/to/scan
docker run --rm -v $(pwd):/scan-target:ro bugzora:latest fs /scan-target
```

### Memory issues
```bash
docker run --rm --memory=2g bugzora:latest image large-image:latest
```

## Performance Tips

- Use Trivy cache volume for faster repeated scans
- Use quiet mode for faster output
- Use JSON output for parsing
- Use multi-arch build for cross-platform support

## Integration Examples

### CI/CD Pipeline
```yaml
- name: Security Scan
  run: |
    docker-compose run --rm bugzora image ${{ github.repository }}:${{ github.sha }} --output json > security-report.json
```

### Automated Scanning
```bash
for image in $(ls images/); do
  docker-compose run --rm bugzora image $image --output json > reports/$image.json
done
```

## Best Practices

1. Always use read-only volumes for filesystem scanning
2. Run as non-root user (already configured)
3. Use specific image tags instead of `latest`
4. Clean up containers after use
5. Monitor resource usage for large scans
6. Use quiet mode in automated environments
7. Cache Trivy database for better performance

## Further Reading

- [DOCKER_OPTIMIZATION.md](DOCKER_OPTIMIZATION.md): Full details on all Docker optimizations, security, and performance features
- [build-docker.sh](build-docker.sh): Multi-arch build script
- [docker-security-scan.sh](docker-security-scan.sh): Automated security scan script

---

**BugZora Docker Guide**  
For more information, see the main [README.md](README.md) and [DOCKER_OPTIMIZATION.md](DOCKER_OPTIMIZATION.md) files. 