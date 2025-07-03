# BugZora v2.1.0 Sürüm Notları

## 🎉 Demo Modu Aktif Edildi

Bu sürüm ile BugZora'nın tam özellikli demo modu aktif edilmiştir.

### 🚀 Yeni Özellikler

#### Demo Modu
- **Demo Modu**: Tüm özellikler demo modunda kullanılabilir
- **Gerçek Tarama**: Demo modunda gerçek Trivy taraması yapılır
- **Gizli Detaylar**: Detaylı sonuçlar gizlenir, sadece bulgu sayıları ve severity seviyeleri gösterilir
- **Demo Mesajları**: Tüm çıktılarda demo modu uyarıları

#### CLI Komutları
- `bugzora image` - Container image tarama (Demo)
- `bugzora fs` - Dosya sistemi tarama (Demo)
- `bugzora repo` - Git repository tarama (Demo)
- `bugzora secret` - Secret tarama (Demo)
- `bugzora license` - Lisans uyumluluk tarama (Demo)

#### Çıktı Formatları
- **Table**: Renkli terminal çıktısı
- **JSON**: Trivy JSON formatı
- **PDF**: Detaylı raporlar
- **SBOM**: Software Bill of Materials
- **CycloneDX**: Standart SBOM formatı
- **SPDX**: Software Package Data Exchange

### 🔧 İyileştirmeler

#### Performans
- Hızlı tarama simülasyonu
- Optimize edilmiş çıktı formatları
- Geliştirilmiş tablo hizalaması

#### Güvenlik
- Trivy entegrasyonu
- Zafiyet tarama
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
- Zafiyet detayları gizlenir
- CVE bilgileri gizlenir
- Paket detayları gizlenir

#### Demo Mesajları
- "🚨 DEMO MODE" uyarıları
- "📧 Contact: license@bugzora.com" bilgisi
- "🔗 For full features: https://bugzora.com/license" linki

### 📋 Kurulum

#### Binary Kurulum
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
# BugZora_Windows_x86_64.tar.gz dosyasını indirip açın
```

#### Paket Kurulum
```bash
# Debian/Ubuntu
sudo dpkg -i BugZora_v2.1.0_linux_amd64.deb

# Red Hat/CentOS
sudo rpm -i BugZora_v2.1.0_linux_amd64.rpm
```

### 🚀 Kullanım Örnekleri

#### Container Image Tarama
```bash
bugzora image alpine:latest
bugzora image nginx:latest --format json
bugzora image ubuntu:20.04 --severity HIGH,CRITICAL
```

#### Dosya Sistemi Tarama
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

#### Lisans Uyumluluk
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

📄 NOT: Bu bir demo sonucudur ve gerçek verileri yansıtır.
🔗 Tüm özellikler için: https://bugzora.com/license
```

### 📦 Sürüm Dosyaları

#### Binary Dosyalar
- `BugZora_Linux_x86_64.tar.gz` - Linux x86_64
- `BugZora_Linux_arm64.tar.gz` - Linux ARM64
- `BugZora_Darwin_x86_64.tar.gz` - macOS x86_64
- `BugZora_Darwin_arm64.tar.gz` - macOS ARM64
- `BugZora_Windows_x86_64.tar.gz` - Windows x86_64
- `BugZora_Windows_arm64.tar.gz` - Windows ARM64

#### Paket Dosyaları
- `BugZora_v2.1.0_linux_amd64.deb` - Debian/Ubuntu x86_64
- `BugZora_v2.1.0_linux_arm64.deb` - Debian/Ubuntu ARM64
- `BugZora_v2.1.0_linux_amd64.rpm` - Red Hat/CentOS x86_64
- `BugZora_v2.1.0_linux_arm64.rpm` - Red Hat/CentOS ARM64

### 📞 İletişim

- **Web Sitesi**: https://bugzora.com
- **E-posta**: license@bugzora.com
- **GitHub**: https://github.com/naimalpermuhacir/BugZora

### 📄 Lisans

Bu demo versiyonu MIT lisansı ile dağıtılmaktadır. Tam özellikli sürüm için lisans satın alınması gerekmektedir.

---

**Sürüm Tarihi**: 2025-01-03  
**Versiyon**: v2.1.0  
**Durum**: Demo Modu 