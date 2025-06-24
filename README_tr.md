# BugZora - Güvenlik Tarama Uygulaması

Copyright © 2025 BugZora <bugzora@bugzora.dev>

Bu uygulama, konteyner imajları ve dosya sistemlerini tarayarak güvenlik açıklarını tespit eden gelişmiş bir güvenlik tarama aracıdır. Trivy altyapısını kullanarak kapsamlı zafiyet analizi yapar ve sonuçları farklı formatlarda sunar.

## 🚀 Özellikler

- **Çoklu Format Desteği**: JSON, PDF ve tablo formatlarında rapor oluşturma
- **Konteyner İmaj Taraması**: Docker Hub ve diğer registry'lerden imaj tarama
- **Dosya Sistemi Taraması**: Yerel dosya sistemlerini tarama
- **İşletim Sistemi Tespiti**: Otomatik OS tespiti ve uygun referans linkleri
- **Çoklu Referans Sistemi**: Her zafiyet için kapsamlı referans linkleri
- **Renkli Terminal Çıktısı**: Okunabilir ve profesyonel tablo formatı
- **Detaylı Raporlama**: Zafiyet istatistikleri ve metadata
- **Yeni Stil Rapor Özeti**: Terminal çıktısının başında hızlı genel bakış için özet tablo.
- **Kalın Tablo Başlıkları & Özet**: Tablo başlıkları ve özet satırları daha okunaklı olması için kalın.
- **Tablolar Arası Ekstra Boşluk**: Terminalde farklı tablo geçişleri daha belirgin.
- **Açıklamalı Legend Alanı**: Tablo sembollerinin anlamı için açıklama.

## 📋 Gereksinimler

- Trivy CLI aracı (kurulum scriptleri tarafından otomatik kurulur)
- İnternet bağlantısı (veritabanı güncellemeleri için)

## 🛠️ Kurulum

### Hızlı Kurulum (Önerilen)

#### Linux & macOS
```bash
# Kurulum scriptini indir ve çalıştır
curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash

# Veya önce indir, sonra çalıştır
wget https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh
chmod +x install.sh
./install.sh
```

#### Windows
```cmd
# PowerShell kullanarak (önerilen)
powershell -ExecutionPolicy Bypass -Command "Invoke-Expression (Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.ps1').Content"

# Veya manuel olarak indir ve çalıştır
# 1. install.ps1 dosyasını indir
# 2. Sağ tıkla "PowerShell ile Çalıştır"
```

```batch
# Command Prompt kullanarak
# install.bat dosyasını indir ve çift tıklayarak çalıştır
```

### Manuel Kurulum

#### Ön Gereksinimler
1. **Trivy Kurulumu**:
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
   
   # Diğer Linux
   # Bkz: https://aquasecurity.github.io/trivy/latest/getting-started/installation/
   ```

2. **BugZora Kurulumu**:
   ```bash
   git clone https://github.com/naimalpermuhacir/BugZora.git
   cd BugZora
   go mod download
   go build -o bugzora .
   ```

### Platform Özel Notları

#### macOS
- **M1/M2 Mac**: ARM64 build'leri otomatik tespit edilir ve kurulur
- **Intel Mac**: x86_64 build'leri kullanılır
- **Homebrew**: Mevcutsa Trivy otomatik olarak Homebrew ile kurulur

#### Linux
- **Ubuntu/Debian, Alpine, Fedora, CentOS, RHEL**: Trivy script tarafından resmi depolardan otomatik kurulur
- **Diğer dağıtımlar**: Trivy manuel kurulum gerekebilir (bkz: Trivy dökümantasyonu)
- **ARM64 desteği**: ARM64 mimarileri için tam destek

#### Windows
- **PowerShell**: Önerilen kurulum yöntemi
- **Command Prompt**: Alternatif batch script mevcut
- **Yönetici hakları**: PATH değişiklikleri için gerekebilir
- **Antivirüs**: Executable'ı işaretleyebilir; gerekirse istisnalara ekleyin

## 🎯 Kullanım

### Hızlı Başlangıç

```bash
# Mevcut komutları kontrol et
bugzora --help

# Konteyner imajı tara (tablo çıktısı)
bugzora image alpine:latest

# Özel registry'den tara
bugzora image registry.example.com/myapp:v1.0.0

# Sessiz modda tara
bugzora image nginx:alpine -q

# Dosya sistemi tara
bugzora fs ./my-application

# Dosya sistemini sessiz modda tara
bugzora fs /path/to/filesystem -q
```

### Konteyner İmaj Taraması

```bash
# Tablo formatında çıktı (varsayılan)
bugzora image ubuntu:20.04

# JSON formatında çıktı
bugzora image ubuntu:20.04 --output json

# PDF formatında çıktı
bugzora image ubuntu:20.04 --output pdf

# Sessiz mod
bugzora image ubuntu:20.04 --quiet
```

### Dosya Sistemi Taraması

```bash
# Tablo formatında çıktı
bugzora fs /path/to/filesystem

# JSON formatında çıktı
bugzora fs /path/to/filesystem --output json

# PDF formatında çıktı
bugzora fs /path/to/filesystem --output pdf
```

## 📊 Çıktı Formatları

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

## 🔗 Referans Sistemi

Her zafiyet için aşağıdaki referans türleri otomatik olarak oluşturulur:

### OS-Specific Referanslar
- **Ubuntu**: Ubuntu Security, Ubuntu Tracker
- **Debian**: Debian Security Tracker, Debian Security
- **Alpine**: Alpine Security
- **Red Hat**: Red Hat Security, Red Hat Bugzilla

### Genel CVE Referansları
- **AquaSec**: Birincil zafiyet analizi
- **CVE Details**: Kapsamlı CVE bilgileri
- **MITRE**: Resmi CVE veritabanı
- **NVD**: National Vulnerability Database

## 📁 Çıktı Dosyaları

Raporlar aşağıdaki isimlendirme kuralıyla oluşturulur:
- `report-{target}.json` - JSON raporu
- `report-{target}.pdf` - PDF raporu

Örnek:
- `report-ubuntu-20.04.json`
- `report-ubuntu-20.04.pdf`
- `report-test-fs.json`

## 🎨 Örnek Çıktılar

### Tablo Formatı
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

### JSON Formatı
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

## 🔧 Geliştirme

### Proje Yapısı
```
BugZora/
├── cmd/           # CLI komutları
├── pkg/           # Ana paketler
│   ├── report/    # Raporlama modülü
│   └── vuln/      # Zafiyet tarama modülü
├── db/            # Trivy veritabanı
└── main.go        # Ana uygulama
```

### Bağımlılıklar
- `github.com/spf13/cobra` - CLI framework
- `github.com/aquasecurity/trivy` - Zafiyet tarama motoru
- `github.com/olekukonko/tablewriter` - Tablo oluşturma
- `github.com/fatih/color` - Renkli terminal çıktısı
- `github.com/jung-kurt/gofpdf` - PDF oluşturma

## 🤝 Katkıda Bulunma

1. Fork yapın
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'Add amazing feature'`)
4. Branch'inizi push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## 📄 Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için `LICENSE` dosyasına bakın.

## 🆘 Destek

Sorunlarınız için:
1. GitHub Issues sayfasını kullanın
2. Detaylı hata mesajları ve log'lar ekleyin
3. Kullandığınız işletim sistemi ve versiyonları belirtin

## 🔄 Güncellemeler

- **v1.0.0**: İlk sürüm - temel tarama özellikleri
- **v1.1.0**: Çoklu referans sistemi eklendi
- **v1.2.0**: JSON ve PDF format desteği eklendi
- **v1.3.0**: Gelişmiş raporlama ve metadata eklendi

---

**BugZora** - Güvenlik Tarama Uygulaması  
Copyright © 2025 BugZora <bugzora@bugzora.dev>  
MIT License 