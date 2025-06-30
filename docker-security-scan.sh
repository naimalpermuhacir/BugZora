#!/bin/bash

# BugZora Docker Security Scanning Script
# Scans the BugZora container for vulnerabilities

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
IMAGE_NAME="bugzora"
VERSION=${1:-"latest"}
SCAN_RESULTS_DIR="security-scan-results"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

echo -e "${BLUE}🔒 BugZora Docker Security Scan${NC}"
echo -e "${BLUE}==============================${NC}"
echo -e "Image: ${GREEN}${IMAGE_NAME}:${VERSION}${NC}"
echo -e "Timestamp: ${GREEN}${TIMESTAMP}${NC}"
echo ""

# Create results directory
mkdir -p ${SCAN_RESULTS_DIR}

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker is not running. Please start Docker and try again.${NC}"
    exit 1
fi

# Check if Trivy is available
if ! command -v trivy &> /dev/null; then
    echo -e "${YELLOW}⚠️  Trivy not found. Installing Trivy...${NC}"
    curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin
fi

# Check if image exists
if ! docker image inspect ${IMAGE_NAME}:${VERSION} > /dev/null 2>&1; then
    echo -e "${YELLOW}⚠️  Image ${IMAGE_NAME}:${VERSION} not found. Building...${NC}"
    docker build -t ${IMAGE_NAME}:${VERSION} .
fi

echo -e "${YELLOW}🔍 Starting security scan...${NC}"

# Run Trivy scan
echo -e "${YELLOW}📊 Scanning for vulnerabilities...${NC}"
trivy image \
    --format json \
    --output ${SCAN_RESULTS_DIR}/vulnerabilities_${TIMESTAMP}.json \
    --severity HIGH,CRITICAL \
    ${IMAGE_NAME}:${VERSION}

# Run Trivy scan with table output
echo -e "${YELLOW}📋 Generating table report...${NC}"
trivy image \
    --format table \
    --output ${SCAN_RESULTS_DIR}/vulnerabilities_${TIMESTAMP}.txt \
    --severity HIGH,CRITICAL \
    ${IMAGE_NAME}:${VERSION}

# Run Trivy config scan
echo -e "${YELLOW}⚙️  Scanning configuration...${NC}"
trivy config \
    --format json \
    --output ${SCAN_RESULTS_DIR}/config_${TIMESTAMP}.json \
    .

# Run Trivy config scan with table output
trivy config \
    --format table \
    --output ${SCAN_RESULTS_DIR}/config_${TIMESTAMP}.txt \
    .

# Run Trivy secret scan
echo -e "${YELLOW}🔐 Scanning for secrets...${NC}"
trivy image \
    --security-checks secret \
    --format json \
    --output ${SCAN_RESULTS_DIR}/secrets_${TIMESTAMP}.json \
    ${IMAGE_NAME}:${VERSION}

# Run Trivy secret scan with table output
trivy image \
    --security-checks secret \
    --format table \
    --output ${SCAN_RESULTS_DIR}/secrets_${TIMESTAMP}.txt \
    ${IMAGE_NAME}:${VERSION}

# Generate summary
echo -e "${YELLOW}📝 Generating summary...${NC}"
cat > ${SCAN_RESULTS_DIR}/summary_${TIMESTAMP}.md << EOF
# BugZora Security Scan Summary

**Scan Date:** $(date)
**Image:** ${IMAGE_NAME}:${VERSION}
**Scanner:** Trivy

## Scan Results

### Vulnerability Scan
- **JSON Report:** vulnerabilities_${TIMESTAMP}.json
- **Table Report:** vulnerabilities_${TIMESTAMP}.txt

### Configuration Scan
- **JSON Report:** config_${TIMESTAMP}.json
- **Table Report:** config_${TIMESTAMP}.txt

### Secret Scan
- **JSON Report:** secrets_${TIMESTAMP}.json
- **Table Report:** secrets_${TIMESTAMP}.txt

## Quick Stats

\`\`\`bash
# Vulnerability count by severity
grep -c "CRITICAL\|HIGH" ${SCAN_RESULTS_DIR}/vulnerabilities_${TIMESTAMP}.txt

# Configuration issues
grep -c "FAIL" ${SCAN_RESULTS_DIR}/config_${TIMESTAMP}.txt

# Secrets found
grep -c "SECRET" ${SCAN_RESULTS_DIR}/secrets_${TIMESTAMP}.txt
\`\`\`

## Recommendations

1. Review all HIGH and CRITICAL vulnerabilities
2. Fix configuration issues
3. Remove any hardcoded secrets
4. Update base images regularly
5. Implement automated scanning in CI/CD

EOF

# Check scan results
VULN_COUNT=$(grep -c "CRITICAL\|HIGH" ${SCAN_RESULTS_DIR}/vulnerabilities_${TIMESTAMP}.txt 2>/dev/null || echo "0")
CONFIG_ISSUES=$(grep -c "FAIL" ${SCAN_RESULTS_DIR}/config_${TIMESTAMP}.txt 2>/dev/null || echo "0")
SECRETS_FOUND=$(grep -c "SECRET" ${SCAN_RESULTS_DIR}/secrets_${TIMESTAMP}.txt 2>/dev/null || echo "0")

echo -e "${GREEN}✅ Security scan completed!${NC}"
echo ""
echo -e "${BLUE}📊 Scan Results:${NC}"
echo -e "  Vulnerabilities (HIGH/CRITICAL): ${RED}${VULN_COUNT}${NC}"
echo -e "  Configuration Issues: ${YELLOW}${CONFIG_ISSUES}${NC}"
echo -e "  Secrets Found: ${RED}${SECRETS_FOUND}${NC}"
echo -e "  Results Directory: ${GREEN}${SCAN_RESULTS_DIR}${NC}"
echo ""

# Show critical findings
if [ "$VULN_COUNT" -gt 0 ]; then
    echo -e "${RED}🚨 Critical Findings:${NC}"
    grep -A 2 -B 2 "CRITICAL\|HIGH" ${SCAN_RESULTS_DIR}/vulnerabilities_${TIMESTAMP}.txt | head -20
    echo ""
fi

# Show configuration issues
if [ "$CONFIG_ISSUES" -gt 0 ]; then
    echo -e "${YELLOW}⚠️  Configuration Issues:${NC}"
    grep -A 1 -B 1 "FAIL" ${SCAN_RESULTS_DIR}/config_${TIMESTAMP}.txt | head -10
    echo ""
fi

# Show secrets
if [ "$SECRETS_FOUND" -gt 0 ]; then
    echo -e "${RED}🔐 Secrets Found:${NC}"
    grep -A 1 -B 1 "SECRET" ${SCAN_RESULTS_DIR}/secrets_${TIMESTAMP}.txt | head -10
    echo ""
fi

echo -e "${BLUE}📁 All reports saved to: ${GREEN}${SCAN_RESULTS_DIR}${NC}"
echo -e "${GREEN}🎉 Security scan process completed!${NC}" 