<!-- CI debug adımı testi için dummy değişiklik -->
# BugZora 🔒

A comprehensive security scanner for container images and filesystems, built on top of Trivy with enhanced reporting, policy enforcement, and SBOM generation capabilities.

[![CI/CD Pipeline](https://github.com/naimalpermuhacir/BugZora/workflows/CI%2FCD%20Pipeline/badge.svg)](https://github.com/naimalpermuhacir/BugZora/actions)
[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Trivy](https://img.shields.io/badge/Trivy-0.63.0+-orange.svg)](https://github.com/aquasecurity/trivy)

## ✨ Features

- **Container Image Scanning**: Scan Docker images from any registry
- **Filesystem Analysis**: Security analysis of local filesystems
- **Multiple Output Formats**: Table, JSON, PDF, CycloneDX, SPDX
- **Policy Enforcement**: OPA/Rego-based security policies
- **Comprehensive References**: OS-specific vulnerability links with color-coded output
- **Docker Integration**: Optimized Docker images with multi-arch support
- **SBOM Generation**: Software Bill of Materials in multiple formats (CycloneDX, SPDX)
- **SBOM Analytics**: Advanced analytics and insights from SBOM data
- **SBOM Validation**: Validate SBOM files for compliance
- **SBOM Comparison**: Compare two SBOM files for differences
- **SBOM Merge**: Merge multiple SBOM files into one
- **Enhanced Terminal Output**: Color-coded severity levels and reference links
- **Optimized Table Display**: Clean, aligned tables with shortened reference links

## 🎯 Demo Mode vs Licensed Version

### 🆓 Demo Mode (Free)
**Available Features:**
- ✅ Real vulnerability scanning with Trivy
- ✅ Container image scanning (alpine:latest, ubuntu:20.04, nginx:latest)
- ✅ Filesystem scanning
- ✅ Secret detection
- ✅ License scanning
- ✅ Repository scanning
- ✅ SBOM generation (CycloneDX, SPDX)
- ✅ Color-coded terminal output
- ✅ Basic policy enforcement
- ✅ JSON and table output formats

**Demo Mode Limitations:**
- 🔒 Only shows vulnerability counts and severity levels
- 🔒 Detailed vulnerability information is hidden
- 🔒 Full SBOM analytics are limited
- 🔒 Advanced policy features are restricted
- 🔒 "This is a demo result but reflects real data" message displayed

**Demo Mode Examples:**
```bash
# Demo mode - shows only counts
./bugzora image alpine:latest
# Output: Found 5 vulnerabilities (2 HIGH, 3 MEDIUM)
#         This is a demo result but reflects real data

./bugzora fs /path/to/filesystem
# Output: Found 12 secrets, 3 license issues
#         This is a demo result but reflects real data
```

### 🔐 Licensed Version (Premium)
**Full Features:**
- ✅ Complete vulnerability details with CVE information
- ✅ Full SBOM analytics and insights
- ✅ Advanced policy enforcement with custom rules
- ✅ Comprehensive reporting with PDF generation
- ✅ SBOM comparison and merge capabilities
- ✅ Detailed reference links and remediation guidance
- ✅ Enterprise-grade security compliance
- ✅ Priority support and updates
- ✅ Custom integrations and APIs

**Licensed Version Examples:**
```bash
# Licensed version - full details
./bugzora image alpine:latest
# Output: Complete vulnerability table with CVE details,
#         reference links, and remediation guidance

./bugzora analytics sbom.json
# Output: Comprehensive SBOM analytics with dependency graphs,
#         risk scoring, and trend analysis
```

## 🚀 Quick Start

### Installation

```bash
# Download and install
curl -sSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/main/install.sh | bash

# Or build from source
git clone https://github.com/naimalpermuhacir/BugZora.git
cd BugZora
go build -o bugzora .
```

### Basic Usage

```bash
# Scan a container image
./bugzora image alpine:latest

# Scan a filesystem
./bugzora fs /path/to/filesystem

# Generate SBOM
./bugzora sbom alpine:latest --format cyclonedx

# Analyze SBOM
./bugzora analytics sbom-file.json

# Compare SBOMs
./bugzora diff sbom1.json sbom2.json
```

## 📊 Enhanced Output Features

### Color-Coded Severity Levels
- **CRITICAL**: Red (bold)
- **HIGH**: Red
- **MEDIUM**: Yellow
- **LOW**: Cyan
- **UNKNOWN**: White

### Reference Links
- OS-specific vulnerability links (Ubuntu, Debian, Alpine, Red Hat)
- CVE database references (NVD, MITRE, CVE Details)
- Color-coded reference links for better readability
- Shortened link format to prevent table misalignment

## 🔧 Advanced Features

### Policy Enforcement
```bash
# Use custom policy file
./bugzora image nginx:latest --policy-file policy.yaml
```

### SBOM Operations
```bash
# Generate SBOM in different formats
./bugzora sbom alpine:latest --format cyclonedx --output sbom.cdx
./bugzora sbom alpine:latest --format spdx --output sbom.spdx

# Validate SBOM
./bugzora validate sbom.cdx

# Merge multiple SBOMs
./bugzora merge sbom1.json sbom2.json --output merged.json

# Compare SBOMs
./bugzora diff sbom1.json sbom2.json
```

### Analytics
```bash
# Generate comprehensive analytics
./bugzora analytics sbom.json --output analytics.json
```

## 📋 Output Formats

- **Table**: Enhanced terminal output with color coding
- **JSON**: Structured vulnerability data
- **PDF**: Comprehensive security reports
- **CycloneDX**: Standard SBOM format
- **SPDX**: Software Package Data Exchange format

## 🐳 Docker Support

```bash
# Run with Docker
docker run --rm -v $(pwd):/workspace naimalpermuhacir/bugzora:latest image alpine:latest

# Build optimized image
docker build -t bugzora .
```

## 🔍 Reference Integration

BugZora provides comprehensive reference links for each vulnerability:
- **Primary**: Direct vulnerability information
- **OS-Specific**: Ubuntu, Debian, Alpine, Red Hat advisories
- **CVE Databases**: NVD, MITRE, CVE Details
- **Vendor Advisories**: Official security bulletins

## 📈 Recent Improvements

- ✅ Enhanced terminal output with color coding
- ✅ Optimized table display with proper alignment
- ✅ Shortened reference links to prevent table misalignment
- ✅ Comprehensive SBOM support (CycloneDX, SPDX)
- ✅ Advanced SBOM analytics and comparison
- ✅ Policy enforcement with OPA/Rego
- ✅ Multi-format output support
- ✅ **CI/CD Pipeline Optimization**: Disk space management and cache optimization
- ✅ **Code Quality**: Exported function and type documentation compliance
- ✅ **Lint Compliance**: golint format compliance for all exported items
- ✅ **Test Organization**: Centralized test artifacts in test-artifacts/ directory
- ✅ **Git Ignore**: Proper test output file management
- ✅ **Demo Mode**: Real scanning with limited output for free users
- ✅ **Licensed Features**: Full functionality for premium users

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Documentation](https://bugzora.dev/docs)
- [Issues](https://github.com/naimalpermuhacir/BugZora/issues)
- [Discussions](https://github.com/naimalpermuhacir/BugZora/discussions) 