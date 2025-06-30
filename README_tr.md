# BugZora 🔒

Container image'ları ve dosya sistemleri için kapsamlı güvenlik tarayıcısı. Trivy üzerine inşa edilmiş, gelişmiş raporlama ve policy enforcement özellikleri ile.

## ✨ Özellikler

- **Container Image Tarama**: Herhangi bir registry'den Docker image'larını tara
- **Dosya Sistemi Analizi**: Yerel dosya sistemlerinin güvenlik analizi
- **Çoklu Çıktı Formatları**: Tablo, JSON, PDF, CycloneDX, SPDX
- **Policy Enforcement**: OPA/Rego tabanlı güvenlik politikaları
- **Kapsamlı Referanslar**: İşletim sistemi özel zafiyet linkleri
- **Docker Entegrasyonu**: Çoklu mimari desteği ile optimize edilmiş Docker image'ları
- **SBOM Üretimi**: Çoklu formatta Yazılım Malzeme Listesi
- **Tam Trivy CLI Desteği**: Tüm Trivy parametreleri ve seçenekleri

## 🚀 Hızlı Başlangıç

### Kurulum

```bash
# En son sürümü indir
curl -L https://github.com/naimalpermuhacir/BugZora/releases/latest/download/bugzora_$(uname -s)_$(uname -m).tar.gz | tar -xz
sudo mv bugzora /usr/local/bin/

# Veya Docker kullan
docker pull naimalpermuhacir/bugzora:latest
```

### Temel Kullanım

```bash
# Container image tara
bugzora image alpine:latest

# Dosya sistemi tara
bugzora fs /path/to/filesystem

# JSON raporu oluştur
bugzora image nginx:latest --format json

# SBOM oluştur
bugzora image ubuntu:20.04 --format cyclonedx
```

## 📋 Gereksinimler

- **Go 1.21+** (geliştirme için)
- **Trivy CLI** (Docker'da otomatik kurulur)
- **Docker** (isteğe bağlı, containerized kullanım için)

## 🔧 Gelişmiş Kullanım

### Policy Enforcement

```bash
# Varsayılan policy oluştur
bugzora policy create policy.yaml

# Policy enforcement ile tara
bugzora image alpine:latest --policy-file policy.yaml
```

### Çoklu Çıktı Formatları

```bash
# JSON raporu
bugzora image nginx:latest --format json --output report.json

# PDF raporu
bugzora image ubuntu:20.04 --format pdf

# CycloneDX SBOM
bugzora fs /app --format cyclonedx

# SPDX SBOM
bugzora image alpine:latest --format spdx
```

### Gelişmiş Tarama Seçenekleri

   ```bash
# Belirli severity'ler ile tara
bugzora image nginx:latest --severity HIGH,CRITICAL

# Düzeltilmemiş zafiyetleri atla
bugzora fs /app --ignore-unfixed

# Tüm paketleri dahil et
bugzora image alpine:latest --list-all-pkgs

# Çevrimdışı tarama
bugzora fs /app --offline-scan
```

## 🐳 Docker Kullanımı

### Hızlı Tarama

```bash
# Image tara
docker run --rm naimalpermuhacir/bugzora:latest image alpine:latest

# Dosya sistemi tara
docker run --rm -v /path:/scan naimalpermuhacir/bugzora:latest fs /scan
```

### Üretim Kullanımı

```bash
# Optimize edilmiş image oluştur
./build-docker.sh

# Güvenlik taraması çalıştır
./docker-security-scan.sh naimalpermuhacir/bugzora:latest
```

## 📊 Çıktı Formatları

### Tablo Çıktısı (Varsayılan)
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

### JSON Çıktısı
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

### SBOM Çıktısı
- **CycloneDX**: Endüstri standardı JSON formatı
- **SPDX**: Uyumluluk için tag-value formatı

## 🔒 Güvenlik Özellikleri

- **Non-root container çalıştırma**
- **Alpine Linux ile minimal saldırı yüzeyi**
- **Daha küçük image'lar için multi-stage build**
- **Container image'larının güvenlik taraması**
- **Policy tabanlı enforcement**
- **Kapsamlı zafiyet referansları**

## 🛠️ Geliştirme

### Proje Yapısı
```
BugZora/
├── cmd/           # CLI komutları
├── pkg/           # Ana paketler
│   ├── report/    # Raporlama modülü
│   ├── vuln/      # Zafiyet tarama modülü
│   └── policy/    # Policy enforcement
├── db/            # Trivy veritabanı
└── main.go        # Ana uygulama
```

### Kaynak Koddan Derleme

```bash
# Repository'yi klonla
git clone https://github.com/naimalpermuhacir/BugZora.git
cd BugZora

# Binary oluştur
go build -o bugzora .

# Testleri çalıştır
go test ./...

# Docker image oluştur
./build-docker.sh
```

## 📚 Dokümantasyon

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
1. [Dokümantasyonu](how_to_use_tr.md) kontrol edin
2. Mevcut [GitHub Issues](https://github.com/naimalpermuhacir/BugZora/issues) arayın
3. Detaylı bilgilerle yeni issue oluşturun

## 🔄 Güncellemeler

- **v1.3.0**: Tam Trivy CLI desteği, SBOM üretimi, policy enforcement
- **v1.2.0**: Docker optimizasyonları, güvenlik sertleştirme, çoklu mimari desteği
- **v1.1.0**: Gelişmiş raporlama, çoklu referans sistemleri
- **v1.0.0**: Temel tarama özellikleri ile ilk sürüm

## 🙏 Teşekkürler

- [Trivy](https://github.com/aquasecurity/trivy) - Altyapı tarama motoru
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Aqua Security](https://www.aquasec.com/) - Zafiyet veritabanı 