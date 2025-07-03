# BugZora 🐛

Container image'ları ve filesystem'ler için gelişmiş SBOM özelliklerine sahip kapsamlı güvenlik tarama aracı.

[![CI/CD Pipeline](https://github.com/naimalpermuhacir/BugZora/workflows/CI%2FCD%20Pipeline/badge.svg)](https://github.com/naimalpermuhacir/BugZora/actions)
[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Trivy](https://img.shields.io/badge/Trivy-0.63.0+-orange.svg)](https://github.com/aquasecurity/trivy)

## 🚀 Özellikler

### 🔍 Güvenlik Tarama
- **Container Image'ları**: Docker image'ları için zafiyet tarama
- **Filesystem'ler**: Yerel filesystem'lerin güvenlik analizi
- **Git Repository'leri**: Kapsamlı repository tarama
- **Secret Tespiti**: Yerleşik secret tarama
- **Lisans Uyumluluğu**: Otomatik lisans tespiti

### 📊 SBOM Oluşturma
- **CycloneDX**: Endüstri standardı SBOM formatı
- **SPDX**: Software Package Data Exchange formatı
- **Bağımlılık Grafikleri**: Görsel bağımlılık ilişkileri
- **Detaylı Metadata**: Kapsamlı paket bilgileri
- **Analitik**: SBOM içgörüleri ve trendler

### 🛡️ Güvenlik Özellikleri
- **Policy Enforcement**: OPA/Rego tabanlı politikalar
- **Risk Skorlama**: Otomatik risk değerlendirmesi
- **Uyumluluk Raporlama**: SOC2, ISO27001 desteği
- **Trend Analizi**: Geçmiş zafiyet takibi
- **Anomali Tespiti**: Akıllı tehdit tespiti

### 🎨 Gelişmiş Çıktı
- **Renkli Severity**: CRITICAL, HIGH, MEDIUM, LOW
- **Çoklu Format**: Tablo, JSON, PDF, SARIF
- **Optimize Tablolar**: Profesyonel hizalama
- **Referans Linkleri**: Daha iyi görünüm için kısaltılmış
- **Cross-platform**: Linux, macOS, Windows desteği

## 🎯 Demo Modu vs Lisanslı Versiyon

### 🆓 Demo Modu (Ücretsiz)
**Mevcut Özellikler:**
- ✅ Trivy ile gerçek zafiyet tarama
- ✅ Container image tarama (alpine:latest, ubuntu:20.04, nginx:latest)
- ✅ Filesystem tarama
- ✅ Secret tespiti
- ✅ Lisans tarama
- ✅ Repository tarama
- ✅ SBOM oluşturma (CycloneDX, SPDX)
- ✅ Renkli terminal çıktısı
- ✅ Temel policy enforcement
- ✅ JSON ve tablo çıktı formatları

**Demo Modu Kısıtlamaları:**
- 🔒 Sadece zafiyet sayıları ve severity seviyeleri gösterilir
- 🔒 Detaylı zafiyet bilgileri gizlenir
- 🔒 Tam SBOM analitikleri sınırlıdır
- 🔒 Gelişmiş policy özellikleri kısıtlıdır
- 🔒 "Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır" mesajı gösterilir

**Demo Modu Örnekleri:**
```bash
# Demo modu - sadece sayıları gösterir
./bugzora image alpine:latest
# Çıktı: 5 zafiyet bulundu (2 HIGH, 3 MEDIUM)
#         Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır

./bugzora fs /path/to/filesystem
# Çıktı: 12 secret, 3 lisans sorunu bulundu
#         Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır
```

### 🔐 Lisanslı Versiyon (Premium)
**Tam Özellikler:**
- ✅ CVE bilgileri ile tam zafiyet detayları
- ✅ Tam SBOM analitikleri ve içgörüler
- ✅ Özel kurallarla gelişmiş policy enforcement
- ✅ PDF oluşturma ile kapsamlı raporlama
- ✅ SBOM karşılaştırma ve birleştirme yetenekleri
- ✅ Detaylı referans linkleri ve düzeltme rehberi
- ✅ Kurumsal seviye güvenlik uyumluluğu
- ✅ Öncelikli destek ve güncellemeler
- ✅ Özel entegrasyonlar ve API'ler

**Lisanslı Versiyon Örnekleri:**
```bash
# Lisanslı versiyon - tam detaylar
./bugzora image alpine:latest
# Çıktı: CVE detayları, referans linkleri ve
#         düzeltme rehberi ile tam zafiyet tablosu

./bugzora analytics sbom.json
# Çıktı: Bağımlılık grafikleri, risk skorlama ve
#         trend analizi ile kapsamlı SBOM analitikleri
```

## 📦 Kurulum

### Hızlı Kurulum

```bash
# İndir ve kur
curl -sSL https://raw.githubusercontent.com/naimalpermuhacir/BugZora/main/install.sh | bash

# Veya wget ile
wget -qO- https://raw.githubusercontent.com/naimalpermuhacir/BugZora/main/install.sh | bash
```

### Manuel Kurulum

```bash
# Repository'yi klonla
git clone https://github.com/naimalpermuhacir/BugZora.git
cd BugZora

# Kaynak koddan derle
go build -o bugzora .

# Global olarak kur
sudo mv bugzora /usr/local/bin/
```

### Docker Kurulumu

```bash
# Image'ı çek
docker pull naimalpermuhacir/bugzora:latest

# Container'ı çalıştır
docker run --rm -v $(pwd):/workspace naimalpermuhacir/bugzora:latest
```

## 🛠️ Kullanım

### Temel Tarama

```bash
# Container image tara
bugzora image alpine:latest

# Filesystem tara
bugzora fs /path/to/filesystem

# Git repository tara
bugzora repo https://github.com/user/repo
```

### SBOM Oluşturma

```bash
# CycloneDX SBOM oluştur
bugzora sbom alpine:latest --format cyclonedx

# SPDX SBOM oluştur
bugzora sbom /path/to/filesystem --format spdx

# Bağımlılık grafiği ile oluştur
bugzora sbom nginx:latest --format cyclonedx --dependency-graph
```

### Gelişmiş Özellikler

```bash
# SBOM'dan analitik oluştur
bugzora analytics sbom-file.json

# İki SBOM dosyasını karşılaştır
bugzora diff sbom1.json sbom2.json

# Birden fazla SBOM dosyasını birleştir
bugzora merge sbom1.json sbom2.json

# SBOM dosyasını doğrula
bugzora validate sbom-file.json
```

### Çıktı Formatları

```bash
# Tablo formatı (varsayılan)
bugzora image alpine:latest

# JSON formatı
bugzora image alpine:latest --format json

# PDF raporu
bugzora image alpine:latest --format pdf --output report.pdf

# SARIF formatı
bugzora image alpine:latest --format sarif --output report.sarif
```

## 📊 Örnek Çıktı

### Zafiyet Tarama
```
┌─────────────┬──────────────┬─────────────┬─────────────┬─────────────────────────────────────┐
│   PAKET     │  ZAFİYET     │  SEVERITY   │   DÜZELTME  │            REFERANSLAR              │
├─────────────┼──────────────┼─────────────┼─────────────┼─────────────────────────────────────┤
│   openssl   │   CVE-2023-  │   HIGH      │   1.1.1w    │ https://cve.mitre.org/cgi-bin/     │
│             │   12345      │             │             │ cvename.cgi?name=CVE-2023-12345     │
└─────────────┴──────────────┴─────────────┴─────────────┴─────────────────────────────────────┘
```

### SBOM Analitik
```
📊 SBOM Analitik Raporu
├── Toplam Paket: 1,234
├── Zafiyetler: 45
│   ├── CRITICAL: 2
│   ├── HIGH: 12
│   ├── MEDIUM: 18
│   └── LOW: 13
├── Lisanslar: 15 benzersiz
└── Risk Skoru: 7.2/10
```

## 🔧 Konfigürasyon

### Policy Dosyası Örneği

```yaml
# policy-example.yaml
rules:
  - name: "Kritik Zafiyetler"
    severity: "CRITICAL"
    action: "FAIL"
    
  - name: "Yüksek Zafiyetler"
    severity: "HIGH"
    action: "WARN"
    max_count: 5
    
  - name: "Lisans Uyumluluğu"
    licenses:
      - "MIT"
      - "Apache-2.0"
    action: "FAIL"
```

### Policy ile Kullanım

```bash
# Policy enforcement uygula
bugzora image alpine:latest --policy-file policy-example.yaml

# Özel policy kuralları
bugzora fs /path/to/filesystem --policy-file custom-policy.yaml
```

## 🏗️ Geliştirme

### Gereksinimler

- Go 1.22 veya üzeri
- Trivy (otomatik yönetilir)
- Git

### Geliştirme Ortamı

```bash
# Repository'yi klonla
git clone https://github.com/naimalpermuhacir/BugZora.git
cd BugZora

# Bağımlılıkları yükle
go mod download

# Testleri çalıştır
go test ./... -v

# Binary oluştur
go build -o bugzora .

# Lint kontrolü
golangci-lint run
```

### Docker Geliştirme

```bash
# Geliştirme image'ı oluştur
docker build -t bugzora-dev .

# Container içinde çalıştır
docker run --rm -v $(pwd):/app bugzora-dev
```

## 📈 Son Geliştirmeler

- ✅ Renkli terminal çıktısı ve severity kodlaması
- ✅ Optimize edilmiş tablo gösterimi ve hizalama
- ✅ Kısaltılmış referans linkleri ile tablo kayması önleme
- ✅ Kapsamlı SBOM desteği (CycloneDX, SPDX)
- ✅ Gelişmiş SBOM analitikleri ve karşılaştırma
- ✅ OPA/Rego ile policy enforcement
- ✅ Çoklu format çıktı desteği
- ✅ **CI/CD Pipeline Optimizasyonu**: Disk alanı yönetimi ve cache optimizasyonu
- ✅ **Kod Kalitesi**: Export edilen fonksiyon ve tip dokümantasyon uyumluluğu
- ✅ **Lint Uyumluluğu**: Tüm export edilen öğeler için golint format uyumluluğu
- ✅ **Test Organizasyonu**: test-artifacts/ dizininde merkezi test dosyaları
- ✅ **Git Ignore**: Uygun test çıktı dosyası yönetimi
- ✅ **Demo Modu**: Ücretsiz kullanıcılar için sınırlı çıktı ile gerçek tarama
- ✅ **Lisanslı Özellikler**: Premium kullanıcılar için tam işlevsellik

## 🤝 Katkıda Bulunma

1. Repository'yi fork edin
2. Feature branch oluşturun
3. Değişikliklerinizi yapın
4. Test ekleyin
5. Pull request gönderin

## 📄 Lisans

Bu proje MIT Lisansı altında lisanslanmıştır - detaylar için [LICENSE](LICENSE) dosyasına bakın.

## 🔗 Bağlantılar

- [Dokümantasyon](https://bugzora.dev/docs)
- [Sorunlar](https://github.com/naimalpermuhacir/BugZora/issues)
- [Tartışmalar](https://github.com/naimalpermuhacir/BugZora/discussions)