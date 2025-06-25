FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bugzora .

FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    wget \
    apt-transport-https \
    gnupg \
    lsb-release \
    ca-certificates \
    && wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | apt-key add - \
    && echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | tee -a /etc/apt/sources.list.d/trivy.list \
    && apt-get update \
    && apt-get install -y trivy \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/bugzora .

RUN useradd -m -u 1000 bugzora \
    && chown -R bugzora:bugzora /app

USER bugzora

ENTRYPOINT ["./bugzora"]

CMD ["--help"] 