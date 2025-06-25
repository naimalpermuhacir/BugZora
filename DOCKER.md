# BugZora Docker Guide

This guide explains how to run BugZora using Docker containers.

## Requirements

- Docker
- Docker Compose (optional, but recommended)

## Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository:**
   ```bash
   git clone https://github.com/naimalpermuhacir/BugZora.git
   cd BugZora
   ```

2. **Build the application:**
   ```bash
   docker-compose build
   ```

3. **View help menu:**
   ```bash
   docker-compose run bugzora --help
   ```

4. **Scan a container image:**
   ```bash
   docker-compose run bugzora image ubuntu:20.04
   ```

5. **Scan a filesystem:**
   ```bash
   docker-compose run bugzora fs /scan-target
   ```

### Using Docker directly

1. **Build the image:**
   ```bash
   docker build -t bugzora:latest .
   ```

2. **Run the application:**
   ```bash
   # Help
   docker run --rm bugzora:latest --help
   
   # Image scanning
   docker run --rm bugzora:latest image ubuntu:20.04
   
   # Filesystem scanning (current directory)
   docker run --rm -v $(pwd):/scan-target:ro bugzora:latest fs /scan-target
   ```

## Features

- **Multi-stage build**: Smaller final image size
- **Non-root user**: Security-first approach
- **Trivy integration**: Automatic Trivy installation
- **Volume mounting**: Host filesystem scanning capability
- **Docker socket access**: Docker image scanning capability
- **Lightweight base**: Ubuntu-based with minimal footprint
- **Security hardened**: Read-only filesystem and security options

## Security Features

- Container runs as non-root user (`bugzora`)
- Read-only volume mounting for filesystem scans
- Security options enabled
- Minimal base image reduces attack surface
- No unnecessary packages installed

## Usage Examples

### Different Output Formats

```bash
# JSON output
docker-compose run bugzora image alpine:latest --output json

# PDF output
docker-compose run bugzora image ubuntu:20.04 --output pdf

# Table output (default)
docker-compose run bugzora image debian:11 --output table

# Quiet mode
docker-compose run bugzora image nginx:alpine --quiet
```

### Filesystem Scanning

```bash
# Scan current directory
docker-compose run bugzora fs /scan-target

# Scan specific path
docker run --rm -v /path/to/scan:/target:ro bugzora:latest fs /target

# Scan with different output formats
docker-compose run bugzora fs /scan-target --output json
docker-compose run bugzora fs /scan-target --output pdf
```

### Private Registry Scanning

```bash
# Scan from private registry
docker-compose run bugzora image registry.example.com/myapp:v1.0.0

# With authentication (if needed)
docker run --rm -v ~/.docker:/root/.docker:ro bugzora:latest image private.registry.com/app:latest
```

### Advanced Usage

```bash
# Interactive shell for debugging
docker-compose run --rm bugzora sh

# Check Trivy version
docker-compose run bugzora sh -c "trivy --version"

# Update Trivy database
docker-compose run bugzora sh -c "trivy image --download-db-only"
```

## Docker Compose Configuration

The `docker-compose.yml` file includes:

```yaml
version: '3.8'
services:
  bugzora:
    build: .
    volumes:
      - .:/scan-target:ro  # Mount current directory for filesystem scanning
      - /var/run/docker.sock:/var/run/docker.sock:ro  # Docker socket access
    environment:
      - TRIVY_CACHE_DIR=/tmp/trivy-cache
    working_dir: /app
    user: "1000:1000"  # Non-root user
```

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
# Start Docker Desktop or Docker service
sudo systemctl start docker

# Check Docker status
docker info
```

### Permission errors
```bash
# Add user to docker group
sudo usermod -aG docker $USER

# Log out and back in, or run:
newgrp docker
```

### Trivy cache issues
```bash
# Clear Trivy cache
docker-compose run bugzora sh -c "rm -rf /tmp/trivy-cache/*"

# Rebuild container
docker-compose down
docker-compose build --no-cache
```

### Volume mounting issues
```bash
# Check volume permissions
ls -la /path/to/scan

# Use absolute paths
docker run --rm -v $(pwd):/scan-target:ro bugzora:latest fs /scan-target
```

### Memory issues
```bash
# Increase Docker memory limit
# In Docker Desktop: Settings > Resources > Memory

# Or use Docker with memory limits
docker run --rm --memory=2g bugzora:latest image large-image:latest
```

## Performance Tips

### Optimize for large scans
```bash
# Use quiet mode for faster output
docker-compose run bugzora image large-image:latest --quiet

# Use JSON output for parsing
docker-compose run bugzora image large-image:latest --output json > report.json
```

### Cache Trivy database
```bash
# Create persistent volume for Trivy cache
docker volume create trivy-cache

# Use persistent cache
docker run --rm -v trivy-cache:/tmp/trivy-cache bugzora:latest image ubuntu:20.04
```

## Integration Examples

### CI/CD Pipeline
```yaml
# GitHub Actions example
- name: Security Scan
  run: |
    docker-compose run bugzora image ${{ github.repository }}:${{ github.sha }} --output json > security-report.json
```

### Automated Scanning
```bash
# Scan all images in a directory
for image in $(ls images/); do
  docker-compose run bugzora image $image --output json > reports/$image.json
done
```

## Best Practices

1. **Always use read-only volumes** for filesystem scanning
2. **Run as non-root user** (already configured)
3. **Use specific image tags** instead of `latest`
4. **Clean up containers** after use
5. **Monitor resource usage** for large scans
6. **Use quiet mode** in automated environments
7. **Cache Trivy database** for better performance

## Support

For Docker-related issues:
1. Check the troubleshooting section above
2. Verify Docker and Docker Compose versions
3. Check container logs: `docker-compose logs bugzora`
4. Open an issue on GitHub with detailed error messages

---

**BugZora Docker Guide**  
For more information, see the main [README.md](README.md) file. 