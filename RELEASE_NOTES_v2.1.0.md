# BugZora v2.1.0 Release Notes

## 🎉 Demo Mode Activated

This release enables the full-featured demo mode for BugZora.

### 🚀 New Features

#### Demo Mode
- **Demo Mode**: All features are now available in demo mode
- **Real Scanning**: Real Trivy scanning is performed in demo mode
- **Hidden Details**: Detailed results are hidden, only finding counts and severity levels are shown
- **Demo Messages**: Demo mode warnings in all outputs

#### CLI Commands
- `bugzora image` - Container image scanning (Demo)
- `bugzora fs` - Filesystem scanning (Demo)
- `bugzora repo` - Git repository scanning (Demo)
- `bugzora secret` - Secret scanning (Demo)
- `bugzora license` - License compliance scanning (Demo)

#### Output Formats
- **Table**: Colored terminal output
- **JSON**: Trivy JSON format
- **PDF**: Detailed reports
- **SBOM**: Software Bill of Materials
- **CycloneDX**: Standard SBOM format
- **SPDX**: Software Package Data Exchange

### 🔧 Improvements

#### Performance
- Fast scan simulation
- Optimized output formats
- Improved table alignment

#### Security
- Trivy integration
- Vulnerability scanning
- Policy enforcement
- Secret detection

#### User Experience
- Colored terminal output
- Progress bars
- Demo mode warnings
- English outputs

### 📦 Supported Platforms

#### Binary Files
- **Linux**: x86_64, ARM64
- **Windows**: x86_64, ARM64
- **macOS**: x86_64, ARM64

#### Package Formats
- **Debian/Ubuntu**: .deb packages
- **Red Hat/CentOS**: .rpm packages
- **Universal**: .tar.gz archives

### 🎯 Demo Mode Features

#### Real Scanning
- Real Trivy calls in demo mode
- Original finding counts
- Real severity levels

#### Hidden Details
- Vulnerability details are hidden
- CVE information is hidden
- Package details are hidden

#### Demo Messages
- "🚨 DEMO MODE" warnings
- "📧 Contact: license@bugzora.com" information
- "🔗 For full features: https://bugzora.com/license" link

### 📋 Installation

#### Binary Installation
```bash
# Linux
wget https://github.com/naimalpermuhacir/BugZora/releases/download/v2.1.0/BugZora_Linux_x86_64.tar.gz
tar -xzf BugZora_Linux_x86_64.tar.gz
sudo mv BugZora /usr/local/bin/

# macOS
wget https://github.com/naimalpermuhacir/BugZora/releases/download/v2.1.0/BugZora_Darwin_x86_64.tar.gz
tar -xzf BugZora_Darwin_x86_64.tar.gz
sudo mv BugZora /usr/local/bin/

# Windows
# Download BugZora_Windows_x86_64.tar.gz and extract
```

#### Package Installation
```bash
# Debian/Ubuntu
sudo dpkg -i BugZora_v2.1.0_linux_amd64.deb

# Red Hat/CentOS
sudo rpm -i BugZora_v2.1.0_linux_amd64.rpm
```

### 🚀 Usage Examples

#### Container Image Scanning
```bash
bugzora image alpine:latest
bugzora image nginx:latest --format json
bugzora image ubuntu:20.04 --severity HIGH,CRITICAL
```

#### Filesystem Scanning
```bash
bugzora fs /path/to/filesystem
bugzora fs . --format json --output report.json
bugzora fs /var/www --severity MEDIUM,HIGH
```

#### Repository Scanning
```bash
bugzora repo https://github.com/user/repo
bugzora repo . --format pdf --output report.pdf
```

#### Secret Scanning
```bash
bugzora secret /path/to/code
bugzora secret . --format json
```

#### License Compliance
```bash
bugzora license /path/to/project
bugzora license . --format table
```

### 🔍 Demo Mode Output

```
🚨 DEMO MODE
📧 Contact: license@bugzora.com
──────────────────────────────────────────────────
🔍 Simulating: alpine:latest scanning...
⏳ Scan progress: 100%
✅ Simulation completed!

📊 DEMO RESULTS: alpine:latest
──────────────────────────────────────────────────
PACKAGE         VULNERABILITY   SEVERITY             DESCRIPTION
────────────────────────────────────────────────────────────────────────────────
License Required License Required CRITICAL - 0         License required
License Required License Required HIGH - 0             License required
License Required License Required MEDIUM - 0           License required
License Required License Required LOW - 0              License required

📄 NOTE: This is a demo result but reflects real data.
🔗 For full features: https://bugzora.com/license
```

### 📦 Release Files

#### Binary Files
- `BugZora_Linux_x86_64.tar.gz` - Linux x86_64
- `BugZora_Linux_arm64.tar.gz` - Linux ARM64
- `BugZora_Darwin_x86_64.tar.gz` - macOS x86_64
- `BugZora_Darwin_arm64.tar.gz` - macOS ARM64
- `BugZora_Windows_x86_64.tar.gz` - Windows x86_64
- `BugZora_Windows_arm64.tar.gz` - Windows ARM64

#### Package Files
- `BugZora_v2.1.0_linux_amd64.deb` - Debian/Ubuntu x86_64
- `BugZora_v2.1.0_linux_arm64.deb` - Debian/Ubuntu ARM64
- `BugZora_v2.1.0_linux_amd64.rpm` - Red Hat/CentOS x86_64
- `BugZora_v2.1.0_linux_arm64.rpm` - Red Hat/CentOS ARM64

### 📞 Contact

- **Website**: https://bugzora.com
- **Email**: license@bugzora.com
- **GitHub**: https://github.com/naimalpermuhacir/BugZora

### 📄 License

This demo version is distributed under the MIT license. A license purchase is required for the full-featured version.

---

**Release Date**: 2025-01-03  
**Version**: v2.1.0  
**Status**: Demo Mode 