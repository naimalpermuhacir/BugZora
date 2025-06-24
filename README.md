# BugZora

A comprehensive security scanning tool for container images and filesystems, built with Go.

## Features

- **Container Image Scanning**: Scan Docker images for vulnerabilities
- **Filesystem Scanning**: Scan local filesystems for security issues
- **Multiple Output Formats**: JSON, PDF, and colored table output
- **Cross-Platform Support**: Linux, macOS, and Windows
- **Automated CI/CD**: GitHub Actions integration with security scanning

## 🚀 Features

- **Multiple Format Support**: Generate reports in JSON, PDF, and table formats
- **Container Image Scanning**: Scan images from Docker Hub and other registries
- **Filesystem Scanning**: Scan local filesystems for vulnerabilities
- **OS Detection**: Automatic OS detection and appropriate reference links
- **Multi-Reference System**: Comprehensive reference links for each vulnerability
- **Colored Terminal Output**: Readable and professional table format
- **Detailed Reporting**: Vulnerability statistics and metadata

## 📋 Requirements

- Trivy CLI tool (automatically installed by installation scripts)
- Internet connection (for database updates)

## 🛠️ Installation

### Quick Installation (Recommended)

#### Linux & macOS
```bash
# Download and run the installation script
curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash

# Or download first, then run
wget https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh
chmod +x install.sh
./install.sh
```

#### Windows
```cmd
# Using PowerShell (recommended)
powershell -ExecutionPolicy Bypass -Command "Invoke-Expression (Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.ps1').Content"

# Or download and run manually
# 1. Download install.ps1
# 2. Right-click and "Run with PowerShell"
```

```batch
# Using Command Prompt
# Download install.bat and double-click to run
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
   
   # Other Linux
   # See: https://aquasecurity.github.io/trivy/latest/getting-started/installation/
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
- **M1/M2 Macs**: ARM64 builds are automatically detected and installed
- **Intel Macs**: x86_64 builds are used
- **Homebrew**: Trivy is automatically installed via Homebrew if available

#### Linux
- **Ubuntu/Debian, Alpine, Fedora, CentOS, RHEL**: Trivy is automatically installed by the script using the official repositories
- **Other distributions**: Manual Trivy installation may be required (see Trivy docs)
- **ARM64 support**: Full support for ARM64 architectures

#### Windows
- **PowerShell**: Recommended installation method
- **Command Prompt**: Alternative batch script available
- **Administrator rights**: May be required for PATH modifications
- **Antivirus**: May flag the executable; add to exclusions if needed

## 🎯 Usage

### Quick Start

```bash
# Check available commands
bugzora --help

# Scan a container image (table output)
bugzora image alpine:latest

# Scan from a private registry
bugzora image registry.example.com/myapp:v1.0.0

# Scan with quiet mode
bugzora image nginx:alpine -q

# Scan a filesystem
bugzora fs ./my-application

# Scan filesystem with quiet mode
bugzora fs /path/to/filesystem -q
```

### Container Image Scanning

```bash
# Table format output (default)
bugzora image ubuntu:20.04

# JSON format output
bugzora image ubuntu:20.04 --output json

# PDF format output
bugzora image ubuntu:20.04 --output pdf

# Quiet mode
bugzora image ubuntu:20.04 --quiet
```

### Filesystem Scanning

```bash
# Table format output
bugzora fs /path/to/filesystem

# JSON format output
bugzora fs /path/to/filesystem --output json

# PDF format output
bugzora fs /path/to/filesystem --output pdf
```

## 📊 Output Formats

### 1. Table Format (Default)
Provides colored, readable table output in terminal:
- Vulnerability details
- Multiple reference links
- Colored severity indicators
- Summary statistics

### 2. JSON Format
Generates comprehensive JSON report:
- Scan metadata
- Detailed vulnerability information
- Multiple reference links
- Statistical summary
- Configurable format

### 3. PDF Format
Creates professional PDF report:
- Turkish titles and descriptions
- Colored severity indicators
- Table format vulnerability list
- Reference links
- Summary statistics

## 🔗 Reference System

The following reference types are automatically generated for each vulnerability:

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

## 🔧 Development

### Project Structure
```
guvenlik-app/
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
- `github.com/jung-kurt/gofpdf` - PDF generation

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
- **v1.2.0**: JSON and PDF format support added
- **v1.3.0**: Advanced reporting and metadata added

---

**BugZora** - Security Scanning Application  
Copyright © 2025 BugZora <bugzora@bugzora.dev>  
MIT License 