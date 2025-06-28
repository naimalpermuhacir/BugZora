<!-- CI debug adımı testi için dummy değişiklik -->
# BugZora

BugZora, Trivy motorunu kullanan, konteyner imajları ve dosya sistemleri için gelişmiş bir güvenlik tarama aracıdır.

## Özellikler
- Tüm Trivy CLI parametrelerini destekler (örn. --severity, --scanners, --ignore-unfixed, --exit-code, --skip-dirs, --list-all-pkgs, --offline-scan, --template, --policy, --config, --token, --proxy, --timeout, --download-db-only, --reset, --clear-cache, --debug, --trace, --no-progress, --ignore-policy, --skip-update, --skip-db-update, --skip-policy-update, --security-checks, --compliance, --namespaces, --output, --ignore-ids, --ignore-file, --include-dev-deps, --skip-java-db, --only-update, --refresh, --auto-refresh, --light)
- Çoklu çıktı formatı: table, json, pdf, SARIF, CycloneDX, SPDX
- Policy enforcement (OPA/Rego)
- Secret ve license tarama
- Kubernetes ve repository tarama (yeni sürümlerde)
- Modern, renkli ve özetli tablo raporu
- Multi-arch ve Docker optimizasyonları

## 🚀 Features

- **Container Image Scanning**: Scan Docker images for vulnerabilities
- **Filesystem Scanning**: Scan local filesystems for security issues
- **Multiple Output Formats**: JSON, PDF, and colored table output
- **Cross-Platform Support**: Linux, macOS, and Windows
- **Docker Support**: Containerized deployment with multi-stage builds
- **Multi-Architecture Support**: Build and run on amd64, arm64, arm/v7
- **Automated CI/CD**: GitHub Actions integration with security scanning
- **Modern Report Summary**: A summary table at the top of the terminal output for quick overview
- **Bold Table Headers & Summary**: Table headers and summary lines are bold for better readability
- **Extra Spacing Between Tables**: Visually clear separation between different result tables in terminal output
- **Legend Section**: Explains table symbols for clarity
- **Multi-Reference System**: Comprehensive reference links for each vulnerability
- **Security Hardening**: Non-root user, read-only filesystem, dropped capabilities, health checks

## 🐳 Docker Optimizations & Security

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

## 📋 Requirements

- Trivy CLI tool (automatically installed by installation scripts or Docker)
- Internet connection (for database updates)
- Docker (for containerized usage)

## 🛠️ Installation

### Quick Installation (Recommended)

#### Linux & macOS
```bash
curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash
```

#### Windows
```powershell
powershell -ExecutionPolicy Bypass -Command "Invoke-Expression (Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.ps1').Content"
```

### Docker Installation

#### Using Docker Compose (Recommended)
```bash
git clone https://github.com/naimalpermuhacir/BugZora.git
cd BugZora
docker-compose build
docker-compose run --rm bugzora --help
```

#### Using Docker directly
```bash
docker build -t bugzora:latest .
docker run --rm bugzora:latest --help
```

#### Multi-Architecture Build
```bash
./build-docker.sh v1.2.0
```

#### Security Scan of Container
```bash
./docker-security-scan.sh v1.2.0
ls -la security-scan-results/
```

For detailed Docker usage, see [DOCKER.md](DOCKER.md) and [DOCKER_OPTIMIZATION.md](DOCKER_OPTIMIZATION.md).

### Manual Installation

#### Prerequisites
1. **Install Trivy** (see Trivy docs for your OS)
2. **Install BugZora**:
   ```bash
   git clone https://github.com/naimalpermuhacir/BugZora.git
   cd BugZora
   go mod download
   go build -o bugzora .
   ```

## 🎯 Usage

### Quick Start

```bash
bugzora --help
bugzora image alpine:latest
bugzora fs ./my-application
```

### Docker Usage

```bash
docker-compose run --rm bugzora image ubuntu:20.04
docker-compose run --rm bugzora fs /scan-target
docker run --rm bugzora:latest image ubuntu:20.04
docker run --rm -v $(pwd):/scan-target:ro bugzora:latest fs /scan-target
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

### Output Formats

- Table (default): Colored, readable terminal output
- JSON: Machine-readable, file-based output (requires writable volume)
- PDF: Professional report (requires writable volume)

### Security & Performance

- Non-root user, read-only rootfs, dropped capabilities
- Health checks and resource limits
- Trivy cache volume for fast repeated scans

## 🔧 Development

### Project Structure
```
BugZora/
├── cmd/           # CLI commands
├── pkg/           # Main packages
│   ├── report/    # Reporting module
│   └── vuln/      # Vulnerability scanning module
├── db/            # Trivy database
├── Dockerfile     # Docker container definition
├── docker-compose.yml  # Docker Compose configuration
├── .dockerignore      # Build optimization
├── build-docker.sh    # Multi-arch build script
├── docker-security-scan.sh # Security scan script
├── DOCKER_OPTIMIZATION.md  # Docker optimization guide
└── main.go        # Main application
```

### CI/CD Pipeline
- Test, build, security scan, release (GoReleaser)
- Multi-platform builds
- Automated Docker builds and security scans

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the `LICENSE` file for details.

## 🆘 Support

For issues:
1. Use the GitHub Issues page
2. Include detailed error messages and logs
3. Specify your operating system and versions

## 🔄 Updates

- **v1.0.0**: Initial release - basic scanning features
- **v1.1.0**: Multi-reference system added
- **v1.1.1**: Docker support and CI/CD pipeline
- **v1.2.0**: Docker optimizations, security hardening, multi-arch, advanced reporting
- **v1.3.0**: Advanced reporting and metadata added

---

**BugZora** - Security Scanning Application  
Copyright © 2025 BugZora <bugzora@bugzora.dev>  
MIT License 