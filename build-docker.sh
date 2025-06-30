#!/bin/bash

# BugZora Multi-Architecture Docker Build Script
# Supports: linux/amd64, linux/arm64, linux/arm/v7

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
PLATFORMS="linux/amd64,linux/arm64,linux/arm/v7"
BUILDX_BUILDER="bugzora-builder"

echo -e "${BLUE}🐳 BugZora Multi-Architecture Docker Build${NC}"
echo -e "${BLUE}===========================================${NC}"
echo -e "Image: ${GREEN}${IMAGE_NAME}:${VERSION}${NC}"
echo -e "Platforms: ${GREEN}${PLATFORMS}${NC}"
echo ""

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker is not running. Please start Docker and try again.${NC}"
    exit 1
fi

# Check if buildx is available
if ! docker buildx version > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker buildx is not available. Please install Docker buildx.${NC}"
    exit 1
fi

# Create or use existing builder
echo -e "${YELLOW}🔧 Setting up buildx builder...${NC}"
if ! docker buildx inspect ${BUILDX_BUILDER} > /dev/null 2>&1; then
    echo -e "${YELLOW}Creating new buildx builder: ${BUILDX_BUILDER}${NC}"
    docker buildx create --name ${BUILDX_BUILDER} --use
else
    echo -e "${YELLOW}Using existing buildx builder: ${BUILDX_BUILDER}${NC}"
    docker buildx use ${BUILDX_BUILDER}
fi

# Bootstrap the builder
echo -e "${YELLOW}🚀 Bootstrapping builder...${NC}"
docker buildx inspect --bootstrap

# Build the image
echo -e "${YELLOW}🏗️  Building multi-architecture image...${NC}"
docker buildx build \
    --platform ${PLATFORMS} \
    --tag ${IMAGE_NAME}:${VERSION} \
    --tag ${IMAGE_NAME}:latest \
    --file Dockerfile \
    --push \
    --cache-from type=registry,ref=${IMAGE_NAME}:buildcache \
    --cache-to type=registry,ref=${IMAGE_NAME}:buildcache,mode=max \
    .

# Check build result
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Multi-architecture build completed successfully!${NC}"
    echo ""
    echo -e "${BLUE}📋 Build Summary:${NC}"
    echo -e "  Image: ${GREEN}${IMAGE_NAME}:${VERSION}${NC}"
    echo -e "  Platforms: ${GREEN}${PLATFORMS}${NC}"
    echo -e "  Builder: ${GREEN}${BUILDX_BUILDER}${NC}"
    echo ""
    echo -e "${BLUE}🚀 Usage:${NC}"
    echo -e "  docker run --rm ${IMAGE_NAME}:${VERSION} --help"
    echo -e "  docker run --rm ${IMAGE_NAME}:${VERSION} image ubuntu:20.04"
    echo -e "  docker run --rm -v \$(pwd):/scan-target:ro ${IMAGE_NAME}:${VERSION} fs /scan-target"
else
    echo -e "${RED}❌ Build failed!${NC}"
    exit 1
fi

# Show image info
echo ""
echo -e "${BLUE}📊 Image Information:${NC}"
docker buildx imagetools inspect ${IMAGE_NAME}:${VERSION}

echo ""
echo -e "${GREEN}🎉 Build process completed!${NC}" 