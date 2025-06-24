@echo off
setlocal enabledelayedexpansion

REM BugZora Installation Script for Windows
REM Copyright © 2025 BugZora <bugzora@bugzora.dev>

echo [INFO] Starting BugZora installation for Windows...

REM Check if running as administrator
net session >nul 2>&1
if %errorLevel% == 0 (
    echo [INFO] Running as administrator
) else (
    echo [WARNING] Not running as administrator. Some operations may require elevation.
)

REM Detect architecture
if "%PROCESSOR_ARCHITECTURE%"=="AMD64" (
    set ARCH=x86_64
) else if "%PROCESSOR_ARCHITECTURE%"=="ARM64" (
    set ARCH=arm64
) else (
    echo [ERROR] Unsupported architecture: %PROCESSOR_ARCHITECTURE%
    pause
    exit /b 1
)

echo [INFO] Detected architecture: %ARCH%

REM Get latest version
echo [INFO] Fetching latest version...
for /f "tokens=*" %%i in ('powershell -Command "(Invoke-RestMethod -Uri 'https://api.github.com/repos/naimalpermuhacir/BugZora/releases/latest').tag_name"') do set VERSION=%%i

if "%VERSION%"=="" (
    echo [WARNING] Could not fetch latest version, using v1.3.0
    set VERSION=v1.3.0
)

echo [INFO] Latest version: %VERSION%

REM Create temporary directory
set TEMP_DIR=%TEMP%\bugzora_install_%RANDOM%
mkdir "%TEMP_DIR%"
cd /d "%TEMP_DIR%"

REM Download BugZora
set DOWNLOAD_URL=https://github.com/naimalpermuhacir/BugZora/releases/download/%VERSION%/bugzora_Windows_%ARCH%.zip
echo [INFO] Downloading BugZora from: %DOWNLOAD_URL%

powershell -Command "& {Invoke-WebRequest -Uri '%DOWNLOAD_URL%' -OutFile 'bugzora.zip'}"
if %errorLevel% neq 0 (
    echo [ERROR] Failed to download BugZora
    rmdir /s /q "%TEMP_DIR%"
    pause
    exit /b 1
)

REM Extract files
echo [INFO] Extracting files...
powershell -Command "& {Expand-Archive -Path 'bugzora.zip' -DestinationPath '.' -Force}"

REM Install to Program Files
set INSTALL_DIR=C:\Program Files\BugZora
echo [INFO] Installing to: %INSTALL_DIR%

if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"
copy "bugzora.exe" "%INSTALL_DIR%\"

REM Add to PATH
echo [INFO] Adding BugZora to PATH...
setx PATH "%PATH%;%INSTALL_DIR%" /M
if %errorLevel% neq 0 (
    echo [WARNING] Failed to add to PATH. You may need to add %INSTALL_DIR% manually.
)

REM Install Trivy if not present
where trivy >nul 2>&1
if %errorLevel% neq 0 (
    echo [INFO] Installing Trivy...
    
    REM Download Trivy for Windows
    powershell -Command "& {Invoke-WebRequest -Uri 'https://github.com/aquasecurity/trivy/releases/latest/download/trivy_%VERSION%_Windows-%ARCH%.zip' -OutFile 'trivy.zip'}"
    if %errorLevel% equ 0 (
        powershell -Command "& {Expand-Archive -Path 'trivy.zip' -DestinationPath '%INSTALL_DIR%' -Force}"
        echo [SUCCESS] Trivy installed
    ) else (
        echo [WARNING] Failed to install Trivy automatically
        echo [WARNING] Please install Trivy manually from: https://aquasecurity.github.io/trivy/latest/getting-started/installation/
    )
) else (
    echo [INFO] Trivy is already installed
)

REM Cleanup
cd /d "%USERPROFILE%"
rmdir /s /q "%TEMP_DIR%"

REM Verify installation
echo [INFO] Verifying installation...
"%INSTALL_DIR%\bugzora.exe" --help >nul 2>&1
if %errorLevel% equ 0 (
    echo [SUCCESS] BugZora installation verified!
    echo [INFO] You can now use: bugzora image ^<image^> or bugzora fs ^<path^>
    echo [INFO] Note: You may need to restart your terminal for PATH changes to take effect.
) else (
    echo [ERROR] BugZora installation failed
    pause
    exit /b 1
)

echo [SUCCESS] Installation completed successfully!
pause 