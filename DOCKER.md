# BugZora Docker Guide

Bu rehber, BugZora uygulamasını Docker ile nasıl çalıştıracağınızı açıklar.

## Gereksinimler

- Docker
- Docker Compose (opsiyonel)

## Hızlı Başlangıç

### Docker Compose ile (Önerilen)

1. **Uygulamayı build edin:**
   ```bash
   docker-compose build
   ```

2. **Yardım menüsünü görüntüleyin:**
   ```bash
   docker-compose run bugzora --help
   ```

3. **Container image taraması yapın:**
   ```bash
   docker-compose run bugzora image ubuntu:20.04
   ```

4. **Filesystem taraması yapın:**
   ```bash
   docker-compose run bugzora fs /scan-target
   ```

### Docker ile

1. **Image build edin:**
   ```bash
   docker build -t bugzora:latest .
   ```

2. **Uygulamayı çalıştırın:**
   ```bash
   # Yardım
   docker run --rm bugzora:latest --help
   
   # Image taraması
   docker run --rm bugzora:latest image ubuntu:20.04
   
   # Filesystem taraması (current directory)
   docker run --rm -v $(pwd):/scan-target:ro bugzora:latest fs /scan-target
   ```

## Özellikler

- **Multi-stage build**: Daha küçük final image boyutu
- **Non-root user**: Güvenlik için root olmayan kullanıcı
- **Trivy entegrasyonu**: Otomatik Trivy kurulumu
- **Volume mounting**: Host filesystem'ini tarayabilme
- **Docker socket access**: Docker image'larını tarayabilme

## Güvenlik

- Container non-root kullanıcı ile çalışır
- Read-only volume mounting kullanılır
- Security options aktif
- Minimal base image kullanılır

## Örnekler

### Farklı output formatları:
```bash
# JSON output
docker-compose run bugzora image alpine:latest --output json

# PDF output
docker-compose run bugzora image ubuntu:20.04 --output pdf

# Table output (default)
docker-compose run bugzora image debian:11 --output table
```

### Filesystem taraması:
```bash
# Current directory'yi tara
docker-compose run bugzora fs /scan-target

# Belirli bir path'i tara
docker run --rm -v /path/to/scan:/target:ro bugzora:latest fs /target
```

## Troubleshooting

### Docker daemon hatası:
```bash
# Docker Desktop'ı başlatın veya Docker service'ini başlatın
sudo systemctl start docker
```

### Permission hatası:
```bash
# Docker group'una kullanıcıyı ekleyin
sudo usermod -aG docker $USER
```

### Trivy cache hatası:
```bash
# Container'ı yeniden başlatın
docker-compose down && docker-compose up
``` 