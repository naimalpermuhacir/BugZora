# BugZora

Konteyner imajları, dosya sistemleri ve Git repository'leri için gelişmiş bir güvenlik tarama aracı. BugZora, kapsamlı güvenlik açığı taraması, gizli bilgi tespiti ve lisans uyumluluğu kontrolü sağlar.

## 🚀 Özellikler

- **🔍 Konteyner İmaj Tarama** - Docker imajlarını registry'lerden tara
- **📁 Dosya Sistemi Tarama** - Yerel dosya sistemlerini güvenlik açıkları için tara
- **🔐 Gizli Bilgi Tespiti** - Açığa çıkmış gizli bilgileri, API anahtarlarını ve şifreleri bul
- **📜 Lisans Tarama** - Yazılım lisanslarını ve uyumluluk sorunlarını tespit et
- **🏗️ Repository Tarama** - Git repository'lerini doğrudan tara
- **📊 Çoklu Çıktı Formatı** - Tablo, JSON, PDF, SARIF, CycloneDX, SPDX
- **🎨 Renkli Terminal Çıktısı** - Güzel, detaylı güvenlik açığı tabloları
- **⚡ Politika Uygulaması** - Otomatik güvenlik politikası değerlendirmesi
- **🔧 Tam Trivy CLI Desteği** - Tüm Trivy bayrakları ve seçenekleri mevcut
- **🐳 Docker Desteği** - Hafif konteynerli dağıtım
- **🔒 Güvenlik Odaklı** - Güvenlik en iyi uygulamalarıyla inşa edildi
- **📈 Modern Rapor Özeti** - Hızlı genel bakış için üstte özet tablo
- **🎯 Kalın Tablo Başlıkları ve Özet** - Daha iyi okunabilirlik için kalın başlıklar ve özet satırları
- **📋 Tablolar Arası Ekstra Boşluk** - Farklı sonuç tabloları arasında net ayrım
- **📚 Açıklama Bölümü** - Tablo sembollerini açıklığa kavuşturur
- **🔗 Çoklu Referans Sistemi** - Her zafiyet için kapsamlı referans linkleri
- **🛡️ Güvenlik Sertleştirme** - Root olmayan kullanıcı, salt okunur dosya sistemi, düşürülmüş yetenekler, sağlık kontrolleri

## 📋 Gereksinimler

- Trivy CLI aracı (kurulum scriptleri tarafından otomatik kurulur)
- İnternet bağlantısı (veritabanı güncellemeleri için)
- Docker (konteynerli kullanım için)

## 🛠️ Kurulum

### Hızlı Kurulum (Önerilen)

#### Linux & macOS
```bash
curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash
```

#### Windows
```powershell
powershell -ExecutionPolicy Bypass -Command "Invoke-Expression (Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.ps1').Content"
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

### Docker Kurulumu

```bash
# En son imajı çek
docker pull naimalpermuhacir/bugzora:latest

# Hızlı tarama yap
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest image alpine:latest
```

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

# Secret tarama
bugzora secret ./uygulama

# License tarama
bugzora license ./uygulama

# Repository tarama
bugzora repo https://github.com/user/repo
```

### Temel Komutlar

```bash
# Konteyner imajı tara
bugzora image ubuntu:20.04

# Dosya sistemi tara
bugzora fs ./uygulama

# Gizli bilgi tara
bugzora secret ./uygulama

# Lisans tara
bugzora license ./uygulama

# Git repository tara
bugzora repo https://github.com/user/repo
```

### Gelişmiş Kullanım

```bash
# JSON raporu oluştur
bugzora image alpine:latest --output json

# Belirli önem seviyesiyle tara
bugzora image nginx:latest --severity HIGH,CRITICAL

# Politika uygulaması kullan
bugzora fs ./uygulama --policy-file guvenlik-politikasi.yaml

# Birden fazla tarayıcıyla tara
bugzora image nginx:latest --scanners vuln,secret,license

# Çevrimdışı tarama
bugzora fs ./uygulama --offline-scan

# Konteyner imajı tara (tablo çıktısı)
bugzora image ubuntu:20.04

# Sessiz modda tara
bugzora image nginx:alpine -q

# Dosya sistemini sessiz modda tara
bugzora fs /path/to/filesystem -q

# Secret tarama
bugzora secret ./uygulama

# Repository tarama
bugzora repo https://github.com/user/repo

# Policy enforcement
bugzora fs ./uygulama --policy-file policy-example.yaml
```

## 📊 Çıktı Formatları

BugZora birden fazla çıktı formatını destekler:

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
# Farklı çıktı formatları oluştur
bugzora image alpine:latest --output table
bugzora image alpine:latest --output json
bugzora image alpine:latest --output pdf
bugzora image alpine:latest --output sarif

# Tablo formatında çıktı (varsayılan)
bugzora image ubuntu:20.04

# JSON formatında çıktı
bugzora fs /path/to/filesystem --output json

# PDF formatında çıktı
bugzora fs /path/to/filesystem --output pdf
```

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

## 🔒 Politika Uygulaması

Güvenlik politikaları oluşturarak tarama sonuçlarını otomatik olarak değerlendirin:

### Politika Dosyası Örneği (guvenlik-politikasi.yaml)
```yaml
rules:
  - name: "Kritik Güvenlik Açıkları"
    description: "Kritik güvenlik açığı bulunursa reddet"
    severity: "CRITICAL"
    max_count: 0
    action: "deny"
  - name: "Yüksek Güvenlik Açıkları"
    description: "5'ten fazla yüksek güvenlik açığı bulunursa uyar"
    severity: "HIGH"
    max_count: 5
    action: "warn"
```

### Politikaları Kullanma
```bash
# Politika uygulamasıyla tara
bugzora fs ./uygulama --policy-file guvenlik-politikasi.yaml

# Docker ile politika
docker run --rm -v $(pwd):/scan -v $(pwd)/guvenlik-politikasi.yaml:/scan/politika.yaml \
  naimalpermuhacir/bugzora:latest fs /scan/uygulama --policy-file /scan/politika.yaml
```

## 🐳 Docker Kullanımı

### Temel Docker Komutları

```bash
# Konteyner imajı tara
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest image ubuntu:20.04

# Dosya sistemi tara
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest fs /scan

# Gizli bilgi tara
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest secret /scan

# JSON raporu oluştur
docker run --rm -v $(pwd):/scan naimalpermuhacir/bugzora:latest image alpine:latest --output json > rapor.json
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

## 🐳 Docker Optimizasyonları ve Güvenlik

- **Multi-stage build**: Küçük, production-ready imajlar
- **Alpine Linux base**: Minimal ve güvenli
- **Trivy kurulumu**: En son sürüm, doğrudan GitHub'dan
- **Root olmayan kullanıcı**: Konteyner UID 1000 olarak çalışır
- **Salt okunur root dosya sistemi**: Gelişmiş güvenlik
- **Düşürülmüş yetenekler**: Sadece temel Linux yetenekleri etkin
- **Sağlık kontrolleri**: Dockerfile ve Compose healthcheck desteği
- **Kaynak sınırları**: Compose'da bellek ve CPU sınırları
- **Volume caching**: Daha hızlı taramalar için Trivy cache
- **Uygun etiketler**: OCI ve Docker metadata
- **Multi-arch build script**: amd64, arm64, arm/v7 için `build-docker.sh`
- **Güvenlik tarama scripti**: Otomatik vulnerability/config/secret tarama için `docker-security-scan.sh`

## 🔧 Gelişmiş Yapılandırma

### Tüm Mevcut Bayraklar

BugZora tüm Trivy CLI bayraklarını destekler:

```bash
# Önem seviyesi filtreleme
--severity LOW,MEDIUM,HIGH,CRITICAL

# Tarayıcı seçimi
--scanners vuln,secret,license,config

# Çıktı özelleştirme
--output table,json,pdf,sarif,cyclonedx,spdx

# Tarama seçenekleri
--ignore-unfixed
--skip-dirs node_modules,vendor
--list-all-pkgs
--offline-scan

# Politika seçenekleri
--policy-file politika.yaml
--ignore-policy

# Ağ seçenekleri
--proxy http://proxy:8080
--timeout 5m

# Hata ayıklama seçenekleri
--debug
--trace
--no-progress
```

### Yapılandırma Dosyası

`trivy.yaml` yapılandırma dosyası oluşturun:

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

## 📈 CI/CD Entegrasyonu

### GitHub Actions Örneği

```yaml
name: Güvenlik Taraması
on: [push, pull_request]

jobs:
  security-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: BugZora Güvenlik Taraması Çalıştır
        run: |
          curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash
          bugzora fs . --severity HIGH,CRITICAL --output json > guvenlik-raporu.json
          
      - name: Güvenlik Raporunu Yükle
        uses: actions/upload-artifact@v4
        with:
          name: guvenlik-raporu
          path: guvenlik-raporu.json
```

### GitLab CI Örneği

```yaml
security-scan:
  image: naimalpermuhacir/bugzora:latest
  script:
    - bugzora fs . --severity HIGH,CRITICAL --output json > guvenlik-raporu.json
  artifacts:
    reports:
      security: guvenlik-raporu.json
```

## 🛡️ Güvenlik Özellikleri

- **Root olmayan kullanıcı** - Konteyner UID 1000 olarak çalışır
- **Salt okunur dosya sistemi** - Gelişmiş güvenlik
- **Düşürülmüş yetenekler** - Sadece temel Linux yetenekleri
- **Sağlık kontrolleri** - Docker sağlık izleme
- **Kaynak sınırları** - Bellek ve CPU kısıtlamaları
- **Çoklu mimari desteği** - amd64, arm64, arm/v7

## 📚 Örnekler

### Konteyner İmaj Tarama

```bash
# Temel imaj taraması
bugzora image nginx:latest

# Belirli önem seviyesiyle tara
bugzora image ubuntu:20.04 --severity HIGH,CRITICAL

# Birden fazla tarayıcıyla tara
bugzora image alpine:latest --scanners vuln,secret,license

# Detaylı rapor oluştur
bugzora image nginx:latest --output json --list-all-pkgs
```

### Dosya Sistemi Tarama

```bash
# Mevcut dizini tara
bugzora fs .

# Belirli dizini tara
bugzora fs /yol/uygulama

# Belirli dizinleri atla
bugzora fs . --skip-dirs node_modules,vendor,.git

# Çevrimdışı tarama
bugzora fs . --offline-scan
```

### Gizli Bilgi Tespiti

```bash
# Kodda gizli bilgi tara
bugzora secret ./uygulama

# Belirli kurallarla tara
bugzora secret . --scanners secret

# Gizli bilgi raporu oluştur
bugzora secret . --output json > gizli-bilgi-raporu.json
```

### Lisans Tarama

```bash
# Lisans uyumluluğu için tara
bugzora license ./proje

# Belirli lisansları kontrol et
bugzora license . --scanners license

# Lisans raporu oluştur
bugzora license . --output json > lisans-raporu.json
```

## 🔍 Sorun Giderme

### Yaygın Sorunlar

1. **Trivy bulunamadı**
   ```bash
   # Trivy'yi yeniden kur
   curl -fsSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/master/install.sh | bash
   ```

2. **İzin reddedildi**
   ```bash
   # Çalıştırılabilir yap
   chmod +x bugzora
   ```

3. **Ağ sorunları**
   ```bash
   # Proxy kullan veya çevrimdışı mod
   bugzora image alpine:latest --proxy http://proxy:8080
   bugzora fs . --offline-scan
   ```

### Hata Ayıklama Modu

```bash
# Hata ayıklama çıktısını etkinleştir
bugzora image alpine:latest --debug

# İzleme günlüğünü etkinleştir
bugzora fs . --trace
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

## 📖 Dokümantasyon

- [Kullanım Kılavuzu](how_to_use_tr.md) - Detaylı kullanım talimatları
- [Docker Kılavuzu](DOCKER.md) - Docker kullanımı ve optimizasyonu
- [Proje Durumu](PROJECT_STATE.md) - Mevcut proje durumu

## 🤝 Katkıda Bulunma

1. Repository'yi fork edin
2. Özellik dalı oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi yapın
4. Uygunsa test ekleyin
5. Pull request gönderin

## 📄 Lisans

Bu proje MIT Lisansı altında lisanslanmıştır - detaylar için [LICENSE](LICENSE) dosyasına bakın.

## 🆘 Destek

Sorunlarınız için:
1. GitHub Issues sayfasını kullanın
2. Detaylı hata mesajları ve log'lar ekleyin
3. Kullandığınız işletim sistemi ve versiyonları belirtin

## 🔄 Güncellemeler

- **v1.0.0**: İlk sürüm - temel tarama özellikleri
- **v1.1.0**: Çoklu referans sistemi eklendi
- **v1.1.1**: Docker desteği ve CI/CD pipeline
- **v1.2.0**: Docker optimizasyonları, güvenlik sertleştirme, çoklu mimari, gelişmiş raporlama
- **v1.3.0**: Tam Trivy CLI parametre desteği, policy enforcement, gelişmiş tarama seçenekleri

## 🙏 Teşekkürler

- Aqua Security'nin [Trivy](https://github.com/aquasecurity/trivy) projesi üzerine inşa edildi
- Daha iyi güvenlik tarama araçları ihtiyacından ilham alındı 