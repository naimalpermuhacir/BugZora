<!-- CI debug adımı testi için dummy değişiklik -->
# BugZora

A powerful security scanning tool for container images, filesystems, and Git repositories. BugZora provides comprehensive vulnerability scanning, secret detection, and license compliance checking with beautiful, colored terminal output.

## 🚀 Features

- **🔍 Container Image Scanning** - Scan Docker images from registries
- **📁 Filesystem Scanning** - Scan local filesystems for vulnerabilities  
- **🔐 Secret Detection** - Find exposed secrets, API keys, and passwords
- **📜 License Scanning** - Detect software licenses and compliance issues
- **🏗️ Repository Scanning** - Scan Git repositories directly
- **📊 Multiple Output Formats** - Table, JSON, PDF, SARIF, CycloneDX, SPDX
- **🎨 Colored Terminal Output** - Beautiful, detailed vulnerability tables
- **⚡ Policy Enforcement** - Automated security policy evaluation
- **🔧 Full Trivy CLI Support** - All Trivy flags and options available
- **🐳 Docker Support** - Lightweight containerized deployment
- **🔒 Security Focused** - Built with security best practices
- **📈 Modern Report Summary** - Summary table at the top for quick overview
- **🎯 Bold Table Headers & Summary** - Bold headers and summary lines for better readability
- **📋 Extra Spacing Between Tables** - Clear separation between different result tables
- **📚 Legend Section** - Explains table symbols for clarity
- **🔗 Multi-Reference System** - Comprehensive reference links for each vulnerability
- **🛡️ Security Hardening** - Non-root user, read-only filesystem, dropped capabilities, health checks

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

### Manual Installation

#### Prerequisites
1. **Install Trivy**:
   ```bash
   # macOS
   brew install trivy
   
   # Ubuntu/Debian
   sudo apt-get install wget apt-transport-https gnupg lsb-release
   wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -
   echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | sudo tee -a /etc/apt/sources.list.d/trivy.list
   sudo apt-get update
   sudo apt-get install trivy
   
   # Alpine
   sudo apk update
   sudo apk add --no-cache trivy
   
   # Fedora
   sudo dnf install -y dnf-plugins-core
   sudo dnf config-manager --add-repo https://aquasecurity.github.io/trivy-repo/rpm/releases/fedora/trivy.repo
   sudo dnf install -y trivy
   
   # CentOS/RHEL
   sudo yum install -y yum-utils
   sudo yum-config-manager --add-repo https://aquasecurity.github.io/trivy-repo/rpm/releases/centos/trivy.repo
   sudo yum install -y trivy
   ```

2. **Install BugZora**:
   ```bash
   git clone https://github.com/naimalpermuhacir/BugZora.git
   cd BugZora
   go mod download
   go build -o bugzora .
   ```

### Platform-Specific Notes

#### macOS
- **M1/M2 Mac**: ARM64 builds are automatically detected and installed
- **Intel Mac**: x86_64 builds are used
- **Homebrew**: Trivy is automatically installed via Homebrew if available

#### Linux
- **Ubuntu/Debian, Alpine, Fedora, CentOS, RHEL**: Trivy is automatically installed from official repositories by the script
- **Other distributions**: Manual Trivy installation may be required (see Trivy documentation)
- **ARM64 support**: Full support for ARM64 architectures

#### Windows
- **PowerShell**: Recommended installation method
- **Command Prompt**: Alternative batch script available
- **Administrator rights**: May be required for PATH changes
- **Antivirus**: May flag executable; add to exceptions if needed

### Docker Installation

```bash
# Pull the latest image
docker pull naimalpermuhacir/bugzora:latest

# Run a quick scan
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest image alpine:latest
```

## 🎯 Quick Start

### Basic Commands

```bash
# Scan a container image
bugzora image ubuntu:20.04

# Scan a filesystem
bugzora fs ./my-application

# Scan for secrets
bugzora secret ./my-application

# Scan for licenses
bugzora license ./my-application

# Scan a Git repository
bugzora repo https://github.com/user/repo
```

### Advanced Usage

```bash
# Generate JSON report
bugzora image alpine:latest --output json

# Scan with specific severity levels
bugzora image nginx:latest --severity HIGH,CRITICAL

# Use policy enforcement
bugzora fs ./my-app --policy-file security-policy.yaml

# Scan with multiple scanners
bugzora image nginx:latest --scanners vuln,secret,license

# Offline scanning
bugzora fs ./my-app --offline-scan

# Scan container image (table output)
bugzora image ubuntu:20.04

# Scan in quiet mode
bugzora image nginx:alpine -q

# Scan filesystem in quiet mode
bugzora fs /path/to/filesystem -q

# Secret scanning
bugzora secret ./my-application

# Repository scanning
bugzora repo https://github.com/user/repo

# Policy enforcement
bugzora fs ./my-app --policy-file policy-example.yaml
```

## 📊 Output Formats

BugZora supports multiple output formats:

- **Tablo** (varsayılan) - Güzel renkli terminal çıktısı
- **JSON** - Otomasyon için makine okunabilir format
- **PDF** - Dokümantasyon için profesyonel raporlar
- **SARIF** - Güvenlik araçları için standart format
- **CycloneDX** - Yazılım Malzeme Listesi (SBOM)
- **SPDX** - Yazılım Paketi Veri Değişimi

### 1. Tablo Formatı (Varsayılan)
Terminalde renkli, okunabilir tablo formatında çıktı verir:
- Zafiyet detayları
- Çoklu referans linkleri
- Renkli severity göstergeleri
- Özet istatistikler

### 2. JSON Formatı
Kapsamlı JSON raporu oluşturur:
- Tarama metadata'sı
- Detaylı zafiyet bilgileri
- Çoklu referans linkleri
- İstatistiksel özet
- Yapılandırılabilir format

### 3. PDF Formatı
Profesyonel PDF raporu oluşturur:
- Türkçe başlıklar ve açıklamalar
- Renkli severity göstergeleri
- Tablo formatında zafiyet listesi
- Referans linkleri
- Özet istatistikler

```bash
# Generate different output formats
bugzora image alpine:latest --output table
bugzora image alpine:latest --output json
bugzora image alpine:latest --output pdf
bugzora image alpine:latest --output sarif

# Table format output (default)
bugzora image ubuntu:20.04

# JSON format output
bugzora fs /path/to/filesystem --output json

# PDF format output
bugzora fs /path/to/filesystem --output pdf
```

## 🔗 Reference System

Her zafiyet için aşağıdaki referans türleri otomatik olarak oluşturulur:

### OS-Specific References
- **Ubuntu**: Ubuntu Security, Ubuntu Tracker
- **Debian**: Debian Security Tracker, Debian Security
- **Alpine**: Alpine Security
- **Red Hat**: Red Hat Security, Red Hat Bugzilla

### General CVE References
- **AquaSec**: Primary vulnerability analysis
- **CVE Details**: Comprehensive CVE information
- **MITRE**: Official CVE database
- **NVD**: National Vulnerability Database

## 📁 Output Files
Reports are generated with the following naming convention:
- `report-{target}.json` - JSON report
- `report-{target}.pdf` - PDF report

Examples:
- `report-ubuntu-20.04.json`
- `report-ubuntu-20.04.pdf`
- `report-test-fs.json`

## 🎨 Sample Outputs

### Table Format
```
--- Vulnerability Scan Report for: ubuntu:20.04 ---
+----------+------------------+----------+------------------+------------------+----------------------------+
| PACKAGE  | VULNERABILITY ID | SEVERITY |  INSTALLED VER   |    FIXED VER     |             TITLE          |
+----------+------------------+----------+------------------+------------------+----------------------------+
| libc-bin | CVE-2025-4802    | MEDIUM   | 2.31-0ubuntu9.17 | 2.31-0ubuntu9.18 | glibc: static setuid binary |
|          |                  |          |                  |                  | dlopen may incorrectly search|
|          |                  |          |                  |                  | LD_LIBRARY_PATH             |
+----------+------------------+----------+------------------+------------------+----------------------------+
```

### JSON Format
```json
{
  "scan_info": {
    "scanner": "bugzora",
    "version": "1.0.0",
    "scan_time": "2025-06-24T11:22:10.657964+03:00"
  },
  "summary": {
    "critical": 0,
    "high": 0,
    "medium": 2,
    "low": 0,
    "unknown": 0,
    "total": 2
  },
  "results": [...]
}
```

## 🔒 Policy Enforcement

Create security policies to automatically evaluate scan results:

### Policy File Example (security-policy.yaml)
```yaml
rules:
  - name: "Critical Vulnerabilities"
    description: "Deny if any CRITICAL vulnerabilities are found"
    severity: "CRITICAL"
    max_count: 0
    action: "deny"
  - name: "High Vulnerabilities"
    description: "Warn if more than 5 HIGH vulnerabilities are found"
    severity: "HIGH"
    max_count: 5
    action: "warn"
```

### Using Policies
```bash
# Scan with policy enforcement
bugzora fs ./my-app --policy-file security-policy.yaml

# Docker with policy
docker run --rm -v $(pwd):/scan -v $(pwd)/security-policy.yaml:/scan/policy.yaml \
  naimalpermuhacir/bugzora:latest fs /scan/my-app --policy-file /scan/policy.yaml
```

## 🐳 Docker Usage

### Basic Docker Commands

```bash
# Scan container image
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest image ubuntu:20.04

# Scan filesystem
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest fs /scan

# Scan for secrets
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest secret /scan

# Generate JSON report
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest image alpine:latest --output json > report.json
```

### Docker Compose

```yaml
version: '3.8'
services:
  bugzora:
    image: naimalpermuhacir/bugzora:latest
    volumes:
      - ./:/scan
    working_dir: /scan
    command: image alpine:latest
```

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

## 🔧 Advanced Configuration

### All Available Flags

BugZora supports all Trivy CLI flags:

```bash
# Severity filtering
--severity LOW,MEDIUM,HIGH,CRITICAL

# Scanner selection
--scanners vuln,secret,license,config

# Output customization
--output table,json,pdf,sarif,cyclonedx,spdx

# Scanning options
--ignore-unfixed
--skip-dirs node_modules,vendor
--list-all-pkgs
--offline-scan

# Policy options
--policy-file policy.yaml
--ignore-policy

# Network options
--proxy http://proxy:8080
--timeout 5m

# Debug options
--debug
--trace
--no-progress
```

### Configuration File

Create a `trivy.yaml` configuration file:

```yaml
# trivy.yaml
severity: HIGH,CRITICAL
scanners:
  - vuln
  - secret
  - license
output: table
skip-dirs:
  - node_modules
  - vendor
  - .git
```

## 📈 CI/CD Integration

### GitHub Actions Example

```yaml
name: Security Scan
on: [push, pull_request]

jobs:
  security-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Run BugZora Security Scan
        run: |
          curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash
          bugzora fs . --severity HIGH,CRITICAL --output json > security-report.json
          
      - name: Upload Security Report
        uses: actions/upload-artifact@v4
        with:
          name: security-report
          path: security-report.json
```

### GitLab CI Example

```yaml
security-scan:
  image: naimalpermuhacir/bugzora:latest
  script:
    - bugzora fs . --severity HIGH,CRITICAL --output json > security-report.json
  artifacts:
    reports:
      security: security-report.json
```

## 🛡️ Security Features

- **Non-root user** - Container runs as UID 1000
- **Read-only filesystem** - Enhanced security
- **Dropped capabilities** - Only essential Linux capabilities
- **Health checks** - Docker health monitoring
- **Resource limits** - Memory and CPU constraints
- **Multi-architecture support** - amd64, arm64, arm/v7

## 📚 Examples

### Container Image Scanning

```bash
# Basic image scan
bugzora image nginx:latest

# Scan with specific severity
bugzora image ubuntu:20.04 --severity HIGH,CRITICAL

# Scan with multiple scanners
bugzora image alpine:latest --scanners vuln,secret,license

# Generate detailed report
bugzora image nginx:latest --output json --list-all-pkgs
```

### Filesystem Scanning

```bash
# Scan current directory
bugzora fs .

# Scan specific directory
bugzora fs /path/to/application

# Skip certain directories
bugzora fs . --skip-dirs node_modules,vendor,.git

# Offline scanning
bugzora fs . --offline-scan
```

### Secret Detection

```bash
# Scan for secrets in code
bugzora secret ./my-application

# Scan with specific rules
bugzora secret . --scanners secret

# Generate secret report
bugzora secret . --output json > secrets-report.json
```

### License Scanning

```bash
# Scan for license compliance
bugzora license ./my-project

# Check specific licenses
bugzora license . --scanners license

# Generate license report
bugzora license . --output json > license-report.json
```

## 🔍 Troubleshooting

### Common Issues

1. **Trivy not found**
   ```bash
   # Reinstall Trivy
   curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash
   ```

2. **Permission denied**
   ```bash
   # Make executable
   chmod +x bugzora
   ```

3. **Network issues**
   ```bash
   # Use proxy or offline mode
   bugzora image alpine:latest --proxy http://proxy:8080
   bugzora fs . --offline-scan
   ```

### Debug Mode

```bash
# Enable debug output
bugzora image alpine:latest --debug

# Enable trace logging
bugzora fs . --trace
```

## 🔧 Development

### Project Structure
```
BugZora/
├── cmd/           # CLI commands
├── pkg/           # Main packages
│   ├── report/    # Reporting module
│   └── vuln/      # Vulnerability scanning module
├── db/            # Trivy database
└── main.go        # Main application
```

### Dependencies
- `github.com/spf13/cobra` - CLI framework
- `github.com/aquasecurity/trivy` - Vulnerability scanning engine
- `github.com/olekukonko/tablewriter` - Table creation
- `github.com/fatih/color` - Colored terminal output
- `github.com/jung-kurt/gofpdf` - PDF creation

## 📖 Documentation

- [How to Use Guide](how_to_use.md) - Detailed usage instructions
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

For issues:
1. Use the GitHub Issues page
2. Include detailed error messages and logs
3. Specify your operating system and versions

## 🔄 Updates

- **v1.0.0**: Initial release - basic scanning features
- **v1.1.0**: Multi-reference system added
- **v1.1.1**: Docker support and CI/CD pipeline
- **v1.2.0**: Docker optimizations, security hardening, multi-arch, advanced reporting
- **v1.3.0**: Full Trivy CLI parameter support, policy enforcement, advanced scanning options

## 🙏 Acknowledgments

- Built on top of [Trivy](https://github.com/aquasecurity/trivy) by Aqua Security
- Inspired by the need for better security scanning tools 