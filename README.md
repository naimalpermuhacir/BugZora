<!-- CI debug adımı testi için dummy değişiklik -->
# BugZora 🔒

A comprehensive security scanner for container images and filesystems, built on top of Trivy with enhanced reporting and policy enforcement capabilities.

## ✨ Features

- **Container Image Scanning**: Scan Docker images from any registry
- **Filesystem Analysis**: Security analysis of local filesystems
- **Multiple Output Formats**: Table, JSON, PDF, CycloneDX, SPDX
- **Policy Enforcement**: OPA/Rego-based security policies
- **Comprehensive References**: OS-specific vulnerability links
- **Docker Integration**: Optimized Docker images with multi-arch support
- **SBOM Generation**: Software Bill of Materials in multiple formats
- **Full Trivy CLI Support**: All Trivy parameters and options

## 🚀 Quick Start

### Installation

```bash
# Download latest release
curl -L https://github.com/naimalpermuhacir/BugZora/releases/latest/download/bugzora_$(uname -s)_$(uname -m).tar.gz | tar -xz
sudo mv bugzora /usr/local/bin/

# Or use Docker
docker pull naimalpermuhacir/bugzora:latest
```

### Basic Usage

```bash
# Scan a container image
bugzora image alpine:latest

# Scan a filesystem
bugzora fs /path/to/filesystem

# Generate JSON report
bugzora image nginx:latest --format json

# Generate SBOM
bugzora image ubuntu:20.04 --format cyclonedx
```

## 📋 Requirements

- **Go 1.21+** (for development)
- **Trivy CLI** (automatically installed in Docker)
- **Docker** (optional, for containerized usage)

## 🔧 Advanced Usage

### Policy Enforcement

```bash
# Create default policy
bugzora policy create policy.yaml

# Scan with policy enforcement
bugzora image alpine:latest --policy-file policy.yaml
```

### Multiple Output Formats

   ```bash
# JSON report
bugzora image nginx:latest --format json --output report.json

# PDF report
bugzora image ubuntu:20.04 --format pdf

# CycloneDX SBOM
bugzora fs /app --format cyclonedx

# SPDX SBOM
bugzora image alpine:latest --format spdx
```

### Advanced Scanning Options

```bash
# Scan with specific severities
bugzora image nginx:latest --severity HIGH,CRITICAL

# Skip unfixed vulnerabilities
bugzora fs /app --ignore-unfixed

# Include all packages
bugzora image alpine:latest --list-all-pkgs

# Offline scanning
bugzora fs /app --offline-scan
```

## 🐳 Docker Usage

### Quick Scan

```bash
# Scan image
docker run --rm naimalpermuhacir/bugzora:latest image alpine:latest

# Scan filesystem
docker run --rm -v /path:/scan naimalpermuhacir/bugzora:latest fs /scan
```

### Production Usage

```bash
# Build optimized image
./build-docker.sh

# Run security scan
./docker-security-scan.sh naimalpermuhacir/bugzora:latest
```

## 📊 Output Formats

### Table Output (Default)
```
Report Summary
┌─────────────┬──────┬─────────────────┬─────────┐
│ Target      │ Type │ Vulnerabilities │ Secrets │
├─────────────┼──────┼─────────────────┼─────────┤
│ alpine:3.18 │ os   │ 5               │ -       │
└─────────────┴──────┴─────────────────┴─────────┘

--- Vulnerability Scan Report for: alpine:3.18 ---
🎯 Target: alpine:3.18 (alpine)

┌─────────────┬─────────────────┬──────────┬─────────────────┬─────────────┬─────────────────────────────────────┬─────────────────────────────────────┐
│ Package     │ Vulnerability ID│ Severity │ Installed Ver.  │ Fixed Ver.  │ Title                               │ Reference                           │
├─────────────┼─────────────────┼──────────┼─────────────────┼─────────────┼─────────────────────────────────────┼─────────────────────────────────────┤
│ openssl     │ CVE-2023-5678   │ HIGH     │ 3.0.8-r0        │ 3.0.9-r0    │ OpenSSL vulnerability description   │ 🔍 Primary: https://avd.aquasec.com │
└─────────────┴─────────────────┴──────────┴─────────────────┴─────────────┴─────────────────────────────────────┴─────────────────────────────────────┘
```

### JSON Output
```json
{
  "scan_info": {
    "scanner": "bugzora",
    "version": "1.3.0",
    "scan_time": "2024-01-15T10:30:00Z"
  },
  "results": [
    {
      "target": "alpine:3.18",
      "type": "alpine",
      "vulnerabilities": [...],
  "summary": {
    "critical": 0,
        "high": 2,
        "medium": 3,
    "low": 0,
    "unknown": 0,
        "total": 5
      }
    }
  ]
}
```

### SBOM Output
- **CycloneDX**: Industry-standard JSON format
- **SPDX**: Tag-value format for compliance

## 🔒 Security Features

- **Non-root container execution**
- **Minimal attack surface** with Alpine Linux
- **Multi-stage builds** for smaller images
- **Security scanning** of container images
- **Policy-based enforcement**
- **Comprehensive vulnerability references**

## 🛠️ Development

### Project Structure
```
BugZora/
├── cmd/           # CLI commands
├── pkg/           # Core packages
│   ├── report/    # Reporting module
│   ├── vuln/      # Vulnerability scanning
│   └── policy/    # Policy enforcement
├── db/            # Trivy database
└── main.go        # Application entry
```

### Building from Source

```bash
# Clone repository
git clone https://github.com/naimalpermuhacir/BugZora.git
cd BugZora

# Build binary
go build -o bugzora .

# Run tests
go test ./...

# Build Docker image
./build-docker.sh
```

## 📚 Documentation

- [Usage Guide](how_to_use.md) - Detailed usage instructions
- [Docker Guide](DOCKER.md) - Docker usage and optimization
- [Project State](PROJECT_STATE.md) - Current project status

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For issues and questions:
1. Check the [documentation](how_to_use.md)
2. Search existing [GitHub Issues](https://github.com/naimalpermuhacir/BugZora/issues)
3. Create a new issue with detailed information

## 🔄 Changelog

- **v1.3.0**: Full Trivy CLI support, SBOM generation, policy enforcement
- **v1.2.0**: Docker optimizations, security hardening, multi-arch support
- **v1.1.0**: Enhanced reporting, multiple reference systems
- **v1.0.0**: Initial release with basic scanning capabilities

## 🙏 Acknowledgments

- [Trivy](https://github.com/aquasecurity/trivy) - The underlying scanning engine
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Aqua Security](https://www.aquasec.com/) - Vulnerability database 