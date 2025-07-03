package policy

import (
	"testing"
	"github.com/aquasecurity/trivy/pkg/types"
)

func TestValidatePolicy(t *testing.T) {
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

	// Policy dosyası olmadan test et
	passed, err := ValidatePolicy(report, "")
	if err != nil {
		t.Fatalf("ValidatePolicy failed: %v", err)
	}

	if !passed {
		t.Error("Expected policy validation to pass in demo mode")
	}
}

func TestEnforcePolicy(t *testing.T) {
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

	// Policy enforcement test et
	err := EnforcePolicy(report, "")
	if err != nil {
		t.Fatalf("EnforcePolicy failed: %v", err)
	}

	// Demo modunda her zaman başarılı olmalı
	t.Log("Policy enforcement passed in demo mode")
}
