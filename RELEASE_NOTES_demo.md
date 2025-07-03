# BugZora Demo Release Notes

## 🎉 Demo Modu Release

Bu release, BugZora'nın demo modunda çalışan tam özellikli versiyonunu içerir.

### 🚀 Yeni Özellikler

#### Demo Modu
- **Demo Modu**: Tüm özellikler demo modunda çalışır durumda
- **Gerçek Tarama**: Demo modunda gerçek Trivy taraması yapılır
- **Gizli Detaylar**: Detaylı sonuçlar gizlenir, sadece bulgu sayıları gösterilir
- **Demo Mesajları**: Tüm çıktılarda demo modu uyarıları

#### CLI Komutları
- `bugzora image` - Container image tarama (Demo)
- `bugzora fs` - Filesystem tarama (Demo)
- `bugzora repo` - Git repository tarama (Demo)
- `bugzora secret` - Secret tarama (Demo)
- `bugzora license` - License compliance tarama (Demo)

#### Çıktı Formatları
- **Table**: Renkli terminal çıktısı
- **JSON**: Trivy JSON formatı
- **PDF**: Detaylı raporlar
- **SBOM**: Software Bill of Materials
- **CycloneDX**: Standard SBOM formatı
- **SPDX**: Software Package Data Exchange

### 🔧 Teknik İyileştirmeler

#### Performans
- Hızlı tarama simülasyonu
- Optimize edilmiş çıktı formatları
- Geliştirilmiş tablo hizalaması

#### Güvenlik
- Trivy entegrasyonu
- Vulnerability scanning
- Policy enforcement
- Secret detection

#### Kullanıcı Deneyimi
- Renkli terminal çıktısı
- Progress bar'lar
- Demo modu uyarıları
- İngilizce çıktılar

### 📦 Desteklenen Platformlar

#### Binary Dosyalar
- **Linux**: x86_64, ARM64
- **Windows**: x86_64, ARM64
- **macOS**: x86_64, ARM64

#### Paket Formatları
- **Debian/Ubuntu**: .deb paketleri
- **Red Hat/CentOS**: .rpm paketleri
- **Universal**: .tar.gz arşivleri

### 🎯 Demo Modu Özellikleri

#### Gerçek Tarama
- Demo modunda gerçek Trivy çağrıları
- Orijinal bulgu sayıları
- Gerçek severity seviyeleri

#### Gizli Detaylar
- Vulnerability detayları gizlenir
- CVE bilgileri gizlenir
- Package detayları gizlenir

#### Demo Mesajları
- "🚨 DEMO MODE" uyarıları
- "📧 Contact: license@bugzora.com" bilgisi
- "🔗 For full features: https://bugzora.com/license" linki

### 📋 Kurulum

#### Binary Kurulum
```bash
# Linux
wget https://github.com/naimalpermuhacir/BugZora/releases/download/demo/BugZora_Linux_x86_64.tar.gz
tar -xzf BugZora_Linux_x86_64.tar.gz
sudo mv BugZora /usr/local/bin/

# macOS
wget https://github.com/naimalpermuhacir/BugZora/releases/download/demo/BugZora_Darwin_x86_64.tar.gz
tar -xzf BugZora_Darwin_x86_64.tar.gz
sudo mv BugZora /usr/local/bin/

# Windows
# Download BugZora_Windows_x86_64.tar.gz and extract
```

#### Paket Kurulum
```bash
# Debian/Ubuntu
sudo dpkg -i BugZora_demo_linux_amd64.deb

# Red Hat/CentOS
sudo rpm -i BugZora_demo_linux_amd64.rpm
```

### 🚀 Kullanım Örnekleri

#### Container Image Tarama
```bash
bugzora image alpine:latest
bugzora image nginx:latest --format json
bugzora image ubuntu:20.04 --severity HIGH,CRITICAL
```

#### Filesystem Tarama
```bash
bugzora fs /path/to/filesystem
bugzora fs . --format json --output report.json
bugzora fs /var/www --severity MEDIUM,HIGH
```

#### Repository Tarama
```bash
bugzora repo https://github.com/user/repo
bugzora repo . --format pdf --output report.pdf
```

#### Secret Tarama
```bash
bugzora secret /path/to/code
bugzora secret . --format json
```

#### License Compliance
```bash
bugzora license /path/to/project
bugzora license . --format table
```

### 🔍 Demo Modu Çıktısı

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

📄 NOT: This is a demo result but reflects real data.
🔗 For full features: https://bugzora.com/license
```

### 📦 Release Dosyaları

#### Binary Dosyalar
- `BugZora_Linux_x86_64.tar.gz` (1.9MB) - Linux x86_64
- `BugZora_Linux_arm64.tar.gz` (1.7MB) - Linux ARM64
- `BugZora_Darwin_x86_64.tar.gz` (1.9MB) - macOS x86_64
- `BugZora_Darwin_arm64.tar.gz` (1.8MB) - macOS ARM64
- `BugZora_Windows_x86_64.tar.gz` (2.0MB) - Windows x86_64
- `BugZora_Windows_arm64.tar.gz` (1.8MB) - Windows ARM64

#### Paket Dosyaları
- `BugZora_demo_linux_amd64.deb` (1.9MB) - Debian/Ubuntu x86_64
- `BugZora_demo_linux_arm64.deb` (1.7MB) - Debian/Ubuntu ARM64
- `BugZora_demo_linux_amd64.rpm` (1.9MB) - Red Hat/CentOS x86_64
- `BugZora_demo_linux_arm64.rpm` (1.8MB) - Red Hat/CentOS ARM64

### 📞 İletişim

- **Website**: https://bugzora.com
- **Email**: license@bugzora.com
- **GitHub**: https://github.com/naimalpermuhacir/BugZora

### 📄 Lisans

Bu demo versiyonu MIT lisansı altında dağıtılmaktadır. Tam özellikli versiyon için lisans satın alınması gerekmektedir.

---

**Release Date**: 2025-01-03  
**Version**: demo  
**Status**: Demo Mode 