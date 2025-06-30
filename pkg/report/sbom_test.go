package report

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aquasecurity/trivy/pkg/types"
)

func TestWriteCycloneDX(t *testing.T) {
	// Create test data
	results := types.Results{
		{
			Target: "alpine:latest",
			Type:   "alpine",
			Vulnerabilities: []types.DetectedVulnerability{
				{
					VulnerabilityID:  "CVE-2023-1234",
					PkgName:          "libc-bin",
					InstalledVersion: "2.31-0ubuntu9.17",
					FixedVersion:     "2.31-0ubuntu9.18",
					PrimaryURL:       "https://example.com/cve-2023-1234",
				},
			},
		},
	}

	// Test CycloneDX generation
	err := WriteCycloneDX("test-cyclonedx", results)
	if err != nil {
		t.Fatalf("Failed to generate CycloneDX SBOM: %v", err)
	}

	// Verify file was created
	fileName := "test-cyclonedx-cyclonedx.json"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Fatalf("CycloneDX file was not created: %s", fileName)
	}

	// Verify JSON structure
	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read CycloneDX file: %v", err)
	}

	var bom CycloneDXBOM
	if err := json.Unmarshal(data, &bom); err != nil {
		t.Fatalf("Failed to parse CycloneDX JSON: %v", err)
	}

	// Verify basic structure
	if bom.BOMFormat != "CycloneDX" {
		t.Errorf("Expected BOMFormat 'CycloneDX', got '%s'", bom.BOMFormat)
	}

	if bom.SpecVersion != "1.5" {
		t.Errorf("Expected SpecVersion '1.5', got '%s'", bom.SpecVersion)
	}

	if len(bom.Components) == 0 {
		t.Error("Expected components to be present")
	}

	// Clean up
	os.Remove(fileName)
}

func TestWriteSPDX(t *testing.T) {
	// Create test data
	results := types.Results{
		{
			Target: "ubuntu:20.04",
			Type:   "ubuntu",
			Vulnerabilities: []types.DetectedVulnerability{
				{
					VulnerabilityID:  "CVE-2023-5678",
					PkgName:          "openssl",
					InstalledVersion: "1.1.1f-1ubuntu2",
					FixedVersion:     "1.1.1f-1ubuntu2.1",
					PrimaryURL:       "https://example.com/cve-2023-5678",
				},
			},
		},
	}

	// Test SPDX generation
	err := WriteSPDX("test-spdx", results)
	if err != nil {
		t.Fatalf("Failed to generate SPDX SBOM: %v", err)
	}

	// Verify file was created
	fileName := "test-spdx-spdx.json"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Fatalf("SPDX file was not created: %s", fileName)
	}

	// Verify JSON structure
	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Failed to read SPDX file: %v", err)
	}

	var doc SPDXDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("Failed to parse SPDX JSON: %v", err)
	}

	// Verify basic structure
	if doc.SPDXVersion != "SPDX-2.3" {
		t.Errorf("Expected SPDXVersion 'SPDX-2.3', got '%s'", doc.SPDXVersion)
	}

	if doc.DataLicense != "CC0-1.0" {
		t.Errorf("Expected DataLicense 'CC0-1.0', got '%s'", doc.DataLicense)
	}

	if len(doc.Packages) == 0 {
		t.Error("Expected packages to be present")
	}

	// Clean up
	os.Remove(fileName)
}

func TestGetPackageType(t *testing.T) {
	tests := []struct {
		resultType string
		expected   string
	}{
		{"alpine", "apk"},
		{"debian", "deb"},
		{"ubuntu", "deb"},
		{"redhat", "rpm"},
		{"centos", "rpm"},
		{"amazon", "rpm"},
		{"oracle", "rpm"},
		{"photon", "rpm"},
		{"suse", "rpm"},
		{"cbl-mariner", "rpm"},
		{"wolfi", "apk"},
		{"chainguard", "apk"},
		{"unknown", "generic"},
	}

	for _, test := range tests {
		result := getPackageType(test.resultType)
		if result != test.expected {
			t.Errorf("getPackageType(%s) = %s, expected %s", test.resultType, result, test.expected)
		}
	}
}

func TestGeneratePURL(t *testing.T) {
	tests := []struct {
		pkgName    string
		version    string
		resultType string
		expected   string
	}{
		{"libc-bin", "2.31-0ubuntu9.17", "debian", "pkg:deb/libc-bin@2.31-0ubuntu9.17"},
		{"openssl", "1.1.1f-1ubuntu2", "ubuntu", "pkg:deb/openssl@1.1.1f-1ubuntu2"},
		{"bash", "5.0.17-1", "alpine", "pkg:apk/bash@5.0.17-1"},
		{"test_package", "1.0.0", "redhat", "pkg:rpm/test-package@1.0.0"},
	}

	for _, test := range tests {
		result := generatePURL(test.pkgName, test.version, test.resultType)
		if result != test.expected {
			t.Errorf("generatePURL(%s, %s, %s) = %s, expected %s",
				test.pkgName, test.version, test.resultType, result, test.expected)
		}
	}
}

func TestConvertSeverity(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"CRITICAL", "critical"},
		{"HIGH", "high"},
		{"MEDIUM", "medium"},
		{"LOW", "low"},
		{"UNKNOWN", "unknown"},
		{"critical", "critical"},
		{"high", "high"},
		{"medium", "medium"},
		{"low", "low"},
		{"unknown", "unknown"},
	}

	for _, test := range tests {
		result := convertSeverity(test.input)
		if result != test.expected {
			t.Errorf("convertSeverity(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}
