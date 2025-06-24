# BugZora Installation Script for Windows PowerShell
# Copyright © 2025 BugZora <bugzora@bugzora.dev>

param(
    [switch]$Force,
    [string]$Version
)

# Function to write colored output
function Write-Status {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor Blue
}

function Write-Success {
    param([string]$Message)
    Write-Host "[SUCCESS] $Message" -ForegroundColor Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[WARNING] $Message" -ForegroundColor Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor Red
}

# Check if running as administrator
$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")

if ($isAdmin) {
    Write-Status "Running as administrator"
} else {
    Write-Warning "Not running as administrator. Some operations may require elevation."
}

# Detect architecture
$arch = if ([Environment]::Is64BitOperatingSystem) { "x86_64" } else { "x86" }
if ([Environment]::Is64BitOperatingSystem -and $env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
    $arch = "arm64"
}

Write-Status "Detected architecture: $arch"

# Get latest version if not specified
if (-not $Version) {
    Write-Status "Fetching latest version..."
    try {
        $response = Invoke-RestMethod -Uri "https://api.github.com/repos/naimalpermuhacir/BugZora/releases/latest"
        $Version = $response.tag_name
    } catch {
        Write-Warning "Could not fetch latest version, using v1.3.0"
        $Version = "v1.3.0"
    }
}

Write-Status "Version: $Version"

# Create temporary directory
$tempDir = Join-Path $env:TEMP "bugzora_install_$(Get-Random)"
New-Item -ItemType Directory -Path $tempDir -Force | Out-Null
Set-Location $tempDir

try {
    # Download BugZora
    $downloadUrl = "https://github.com/naimalpermuhacir/BugZora/releases/download/$Version/bugzora_Windows_$arch.zip"
    Write-Status "Downloading BugZora from: $downloadUrl"
    
    Invoke-WebRequest -Uri $downloadUrl -OutFile "bugzora.zip"
    
    # Extract files
    Write-Status "Extracting files..."
    Expand-Archive -Path "bugzora.zip" -DestinationPath "." -Force
    
    # Install to Program Files
    $installDir = "C:\Program Files\BugZora"
    Write-Status "Installing to: $installDir"
    
    if (-not (Test-Path $installDir)) {
        New-Item -ItemType Directory -Path $installDir -Force | Out-Null
    }
    
    Copy-Item "bugzora.exe" $installDir -Force
    
    # Add to PATH
    Write-Status "Adding BugZora to PATH..."
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", "Machine")
    if ($currentPath -notlike "*$installDir*") {
        $newPath = "$currentPath;$installDir"
        [Environment]::SetEnvironmentVariable("PATH", $newPath, "Machine")
        $env:PATH = "$env:PATH;$installDir"
    }
    
    # Install Trivy if not present
    if (-not (Get-Command trivy -ErrorAction SilentlyContinue)) {
        Write-Status "Installing Trivy..."
        try {
            $trivyUrl = "https://github.com/aquasecurity/trivy/releases/latest/download/trivy_${Version}_Windows-$arch.zip"
            Invoke-WebRequest -Uri $trivyUrl -OutFile "trivy.zip"
            Expand-Archive -Path "trivy.zip" -DestinationPath $installDir -Force
            Write-Success "Trivy installed"
        } catch {
            Write-Warning "Failed to install Trivy automatically"
            Write-Warning "Please install Trivy manually from: https://aquasecurity.github.io/trivy/latest/getting-started/installation/"
        }
    } else {
        Write-Status "Trivy is already installed"
    }
    
    # Verify installation
    Write-Status "Verifying installation..."
    $bugzoraPath = Join-Path $installDir "bugzora.exe"
    if (Test-Path $bugzoraPath) {
        & $bugzoraPath --help | Out-Null
        if ($LASTEXITCODE -eq 0) {
            Write-Success "BugZora installation verified!"
            Write-Status "You can now use: bugzora image <image> or bugzora fs <path>"
            Write-Status "Note: You may need to restart your terminal for PATH changes to take effect."
        } else {
            Write-Error "BugZora installation verification failed"
            exit 1
        }
    } else {
        Write-Error "BugZora executable not found"
        exit 1
    }
    
    Write-Success "Installation completed successfully!"
    
} catch {
    Write-Error "Installation failed: $($_.Exception.Message)"
    exit 1
} finally {
    # Cleanup
    Set-Location $env:USERPROFILE
    if (Test-Path $tempDir) {
        Remove-Item $tempDir -Recurse -Force
    }
} 