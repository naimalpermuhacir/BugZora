package policy

import (
	"os"
	"testing"
)

func TestLoadPolicy(t *testing.T) {
	// Create a temporary policy file
	policyContent := `{
		"rules": [
			{
				"name": "Test Rule",
				"description": "Test Description",
				"severity": "CRITICAL",
				"max_count": 0,
				"action": "deny"
			}
		]
	}`

	tmpFile, err := os.CreateTemp("", "policy-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(policyContent); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// Test loading the policy
	policy, err := LoadPolicy(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to load policy: %v", err)
	}

	if len(policy.Rules) != 1 {
		t.Errorf("Expected 1 rule, got %d", len(policy.Rules))
	}

	if policy.Rules[0].Name != "Test Rule" {
		t.Errorf("Expected rule name 'Test Rule', got '%s'", policy.Rules[0].Name)
	}
}

func TestEvaluatePolicy(t *testing.T) {
	policy := &Policy{
		Rules: []Rule{
			{
				Name:     "Critical Vulnerabilities",
				Severity: "CRITICAL",
				MaxCount: 0,
				Action:   "deny",
			},
			{
				Name:     "High Vulnerabilities",
				Severity: "HIGH",
				MaxCount: 5,
				Action:   "warn",
			},
		},
	}

	vulnerabilities := []Vulnerability{
		{
			VulnerabilityID: "CVE-2023-1234",
			Severity:        "CRITICAL",
			PackageName:     "test-package",
		},
		{
			VulnerabilityID: "CVE-2023-5678",
			Severity:        "HIGH",
			PackageName:     "another-package",
		},
	}

	result := EvaluatePolicy(policy, vulnerabilities)

	if result.Passed {
		t.Error("Expected policy to fail due to CRITICAL vulnerability")
	}

	if result.Action != "deny" {
		t.Errorf("Expected action 'deny', got '%s'", result.Action)
	}

	if len(result.Violations) == 0 {
		t.Error("Expected violations to be reported")
	}
}

func TestMatchesRule(t *testing.T) {
	rule := Rule{
		Severity: "HIGH",
		Packages: []string{"openssl", "nginx"},
		VulnIDs:  []string{"CVE-2023-1234"},
	}

	tests := []struct {
		name          string
		vulnerability Vulnerability
		expected      bool
	}{
		{
			name: "Matching vulnerability",
			vulnerability: Vulnerability{
				VulnerabilityID: "CVE-2023-1234",
				Severity:        "HIGH",
				PackageName:     "openssl",
			},
			expected: true,
		},
		{
			name: "Non-matching severity",
			vulnerability: Vulnerability{
				VulnerabilityID: "CVE-2023-1234",
				Severity:        "MEDIUM",
				PackageName:     "openssl",
			},
			expected: false,
		},
		{
			name: "Non-matching package",
			vulnerability: Vulnerability{
				VulnerabilityID: "CVE-2023-1234",
				Severity:        "HIGH",
				PackageName:     "python",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchesRule(rule, tt.vulnerability)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCreateDefaultPolicy(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "default-policy-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	err = CreateDefaultPolicy(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create default policy: %v", err)
	}

	// Load the created policy to verify it
	policy, err := LoadPolicy(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to load created policy: %v", err)
	}

	if len(policy.Rules) != 3 {
		t.Errorf("Expected 3 default rules, got %d", len(policy.Rules))
	}

	// Check for expected rule names
	expectedNames := map[string]bool{
		"Critical Vulnerabilities": false,
		"High Vulnerabilities":     false,
		"Medium Vulnerabilities":   false,
	}

	for _, rule := range policy.Rules {
		if _, exists := expectedNames[rule.Name]; exists {
			expectedNames[rule.Name] = true
		}
	}

	for name, found := range expectedNames {
		if !found {
			t.Errorf("Expected rule '%s' not found in default policy", name)
		}
	}
}

func TestCountVulnerabilitiesBySeverity(t *testing.T) {
	vulnerabilities := []Vulnerability{
		{Severity: "CRITICAL"},
		{Severity: "HIGH"},
		{Severity: "HIGH"},
		{Severity: "MEDIUM"},
		{Severity: "LOW"},
	}

	tests := []struct {
		severity string
		expected int
	}{
		{"CRITICAL", 1},
		{"HIGH", 2},
		{"MEDIUM", 1},
		{"LOW", 1},
		{"UNKNOWN", 0},
	}

	for _, tt := range tests {
		t.Run(tt.severity, func(t *testing.T) {
			count := countVulnerabilitiesBySeverity(vulnerabilities, tt.severity)
			if count != tt.expected {
				t.Errorf("Expected %d %s vulnerabilities, got %d", tt.expected, tt.severity, count)
			}
		})
	}
}
