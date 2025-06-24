#!/bin/bash

# BugZora Installation Script
# Copyright © 2025 BugZora <bugzora@bugzora.dev>

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Detect OS and architecture
detect_system() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case $ARCH in
        x86_64)
            ARCH="x86_64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac
    
    # Convert OS name to proper case for download URL
    case $OS in
        linux)
            OS_NAME="Linux"
            ;;
        darwin)
            OS_NAME="Darwin"
            ;;
        *)
            OS_NAME="Linux"
            ;;
    esac
    
    print_status "Detected system: $OS $ARCH"
}

# Get latest version from GitHub
get_latest_version() {
    print_status "Fetching latest version..."
    VERSION=$(curl -s https://api.github.com/repos/naimalpermuhacir/BugZora/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    
    if [ -z "$VERSION" ]; then
        print_warning "Could not fetch latest version, using v1.3.0"
        VERSION="v1.3.0"
    fi
    
    print_status "Latest version: $VERSION"
}

# Download BugZora
download_bugzora() {
    local download_url="https://github.com/naimalpermuhacir/BugZora/releases/download/$VERSION/bugzora_${OS_NAME}_${ARCH}.tar.gz"
    
    print_status "Downloading BugZora from: $download_url"
    
    # Create temporary directory
    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"
    
    # Download and extract
    if curl -L -o bugzora.tar.gz "$download_url"; then
        tar -xzf bugzora.tar.gz
        print_success "Download completed"
    else
        print_error "Failed to download BugZora"
        rm -rf "$TEMP_DIR"
        exit 1
    fi
}

# Install BugZora
install_bugzora() {
    local install_dir="/usr/local/bin"
    
    # Check if we have write permissions
    if [ ! -w "$install_dir" ]; then
        print_warning "No write permission to $install_dir, trying to use sudo"
        SUDO_CMD="sudo"
    else
        SUDO_CMD=""
    fi
    
    # Install binary
    if [ -f "bugzora" ]; then
        $SUDO_CMD cp bugzora "$install_dir/"
        $SUDO_CMD chmod +x "$install_dir/bugzora"
        print_success "BugZora installed to $install_dir/bugzora"
    else
        print_error "BugZora binary not found in downloaded files"
        exit 1
    fi
}

# Install Trivy if not present
install_trivy() {
    if command -v trivy &> /dev/null; then
        print_status "Trivy is already installed"
        return
    fi
    
    print_status "Installing Trivy..."

    # Dağıtım tespiti
    if [ "$OS" = "linux" ]; then
        if [ -f /etc/os-release ]; then
            . /etc/os-release
            DISTRO=$ID
        else
            DISTRO="unknown"
        fi
    else
        DISTRO="unknown"
    fi

    case $OS in
        darwin)
            if command -v brew &> /dev/null; then
                brew install trivy
                print_success "Trivy installed via Homebrew"
            else
                print_warning "Homebrew not found. Please install Trivy manually:"
                print_warning "Visit: https://aquasecurity.github.io/trivy/latest/getting-started/installation/"
            fi
            ;;
        linux)
            case $DISTRO in
                ubuntu|debian)
                    sudo apt-get update
                    sudo apt-get install -y wget apt-transport-https gnupg lsb-release
                    wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -
                    echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | sudo tee /etc/apt/sources.list.d/trivy.list
                    sudo apt-get update
                    sudo apt-get install -y trivy
                    ;;
                alpine)
                    sudo apk update
                    sudo apk add --no-cache trivy
                    ;;
                fedora)
                    sudo dnf install -y dnf-plugins-core
                    sudo dnf config-manager --add-repo https://aquasecurity.github.io/trivy-repo/rpm/releases/fedora/trivy.repo
                    sudo dnf install -y trivy
                    ;;
                centos|rhel)
                    sudo yum install -y yum-utils
                    sudo yum-config-manager --add-repo https://aquasecurity.github.io/trivy-repo/rpm/releases/centos/trivy.repo
                    sudo yum install -y trivy
                    ;;
                *)
                    # Fallback: Generic install script
                    curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin
                    ;;
            esac
            print_success "Trivy installed"
            ;;
        *)
            print_warning "Please install Trivy manually for your OS"
            print_warning "Visit: https://aquasecurity.github.io/trivy/latest/getting-started/installation/"
            ;;
    esac
}

# Verify installation
verify_installation() {
    if command -v bugzora &> /dev/null; then
        print_success "BugZora installation verified!"
        print_status "Version: $(bugzora --version 2>/dev/null || echo 'v1.3.0')"
        print_status "Try: bugzora --help"
    else
        print_error "BugZora installation failed"
        exit 1
    fi
}

# Cleanup
cleanup() {
    if [ -n "$TEMP_DIR" ] && [ -d "$TEMP_DIR" ]; then
        rm -rf "$TEMP_DIR"
    fi
}

# Main installation process
main() {
    print_status "Starting BugZora installation..."
    
    # Set up cleanup on exit
    trap cleanup EXIT
    
    detect_system
    get_latest_version
    download_bugzora
    install_bugzora
    install_trivy
    verify_installation
    
    print_success "Installation completed successfully!"
    print_status "You can now use: bugzora image <image> or bugzora fs <path>"
}

# Run main function
main "$@" 