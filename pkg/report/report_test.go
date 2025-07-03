package report

import (
	"testing"
	"github.com/aquasecurity/trivy/pkg/types"
)

func TestGenerateTableReport(t *testing.T) {
	// Test report oluştur
	report := types.Report{
		Results: []types.Result{
			{
				Target: "test-target",
				Vulnerabilities: []types.DetectedVulnerability{
					{
						VulnerabilityID: "CVE-2023-1234",
						PkgName:         "test-package",
					},
				},
			},
		},
	}

	// Table report test et
	result, err := GenerateTableReport(report, "table")
	if err != nil {
		t.Fatalf("GenerateTableReport failed: %v", err)
	}

	if result == "" {
		t.Error("Expected non-empty table report")
	}
}

func TestGenerateJSONReport(t *testing.T) {
	// Test report oluştur
	report := types.Report{
		Results: []types.Result{
			{
				Target: "test-target",
				Vulnerabilities: []types.DetectedVulnerability{
					{
						VulnerabilityID: "CVE-2023-1234",
						PkgName:         "test-package",
					},
				},
			},
		},
	}

	// JSON report test et
	result, err := GenerateJSONReport(report)
	if err != nil {
		t.Fatalf("GenerateJSONReport failed: %v", err)
	}

	if result == "" {
		t.Error("Expected non-empty JSON report")
	}
}

func TestGeneratePDFReport(t *testing.T) {
	// Test report oluştur
	report := types.Report{
		Results: []types.Result{
			{
				Target: "test-target",
				Vulnerabilities: []types.DetectedVulnerability{
					{
						VulnerabilityID: "CVE-2023-1234",
						PkgName:         "test-package",
					},
				},
			},
		},
	}

	// PDF report test et
	err := GeneratePDFReport(report, "test-report.pdf")
	if err != nil {
		t.Fatalf("GeneratePDFReport failed: %v", err)
	}
}

func TestGenerateSBOMReport(t *testing.T) {
	// Test report oluştur
	report := types.Report{
		Results: []types.Result{
			{
				Target: "test-target",
				Vulnerabilities: []types.DetectedVulnerability{
					{
						VulnerabilityID: "CVE-2023-1234",
						PkgName:         "test-package",
					},
				},
			},
		},
	}

	// SBOM report test et
	result, err := GenerateSBOMReport(report, "cyclonedx")
	if err != nil {
		t.Fatalf("GenerateSBOMReport failed: %v", err)
	}

	if result == "" {
		t.Error("Expected non-empty SBOM report")
	}
}
