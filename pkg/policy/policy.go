package policy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

// Policy represents a security policy configuration
type Policy struct {
	Rules []Rule `json:"rules" yaml:"rules"`
}

// Rule represents a single policy rule
type Rule struct {
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	Severity    string            `json:"severity" yaml:"severity"`
	MaxCount    int               `json:"max_count" yaml:"max_count"`
	MinCount    int               `json:"min_count" yaml:"min_count"`
	Action      string            `json:"action" yaml:"action"`
	Packages    []string          `json:"packages" yaml:"packages"`
	VulnIDs     []string          `json:"vuln_ids" yaml:"vuln_ids"`
	Conditions  map[string]string `json:"conditions" yaml:"conditions"`
}

// Result represents the result of policy evaluation
type Result struct {
	Passed     bool     `json:"passed"`
	Violations []string `json:"violations"`
	Warnings   []string `json:"warnings"`
	Action     string   `json:"action"`
}

// Vulnerability represents a single vulnerability for policy evaluation
type Vulnerability struct {
	VulnerabilityID string            `json:"vulnerability_id"`
	Severity        string            `json:"severity"`
	PackageName     string            `json:"package_name"`
	PackageVersion  string            `json:"package_version"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	Metadata        map[string]string `json:"metadata"`
}

// LoadPolicy loads a policy from a file
func LoadPolicy(filePath string) (*Policy, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read policy file: %w", err)
	}

	var policy Policy

	// Try YAML first, then JSON
	if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		err = yaml.Unmarshal(data, &policy)
	} else {
		err = json.Unmarshal(data, &policy)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to parse policy file: %w", err)
	}

	return &policy, nil
}

// EvaluatePolicy evaluates vulnerabilities against the policy
func EvaluatePolicy(policy *Policy, vulnerabilities []Vulnerability) *Result {
	result := &Result{
		Passed:     true,
		Violations: []string{},
		Warnings:   []string{},
		Action:     "allow",
	}

	for _, rule := range policy.Rules {
		ruleResult := evaluateRule(rule, vulnerabilities)

		if !ruleResult.passed {
			result.Passed = false
			result.Violations = append(result.Violations, ruleResult.message)

			// Determine action based on rule
			switch strings.ToLower(rule.Action) {
			case "deny", "fail", "stop":
				result.Action = "deny"
			case "warn":
				result.Warnings = append(result.Warnings, ruleResult.message)
			}
		}
	}

	return result
}

type ruleResult struct {
	passed  bool
	message string
}

// evaluateRule evaluates a single rule against vulnerabilities
func evaluateRule(rule Rule, vulnerabilities []Vulnerability) ruleResult {
	// Count vulnerabilities matching the rule
	count := 0
	matchedVulns := []string{}

	for _, vuln := range vulnerabilities {
		if matchesRule(rule, vuln) {
			count++
			matchedVulns = append(matchedVulns, vuln.VulnerabilityID)
		}
	}

	// Check severity-based rules
	if rule.Severity != "" {
		severityCount := countVulnerabilitiesBySeverity(vulnerabilities, rule.Severity)
		if rule.MaxCount >= 0 && severityCount > rule.MaxCount {
			return ruleResult{
				passed: false,
				message: fmt.Sprintf("Rule '%s' violated: Found %d %s vulnerabilities, maximum allowed is %d",
					rule.Name, severityCount, rule.Severity, rule.MaxCount),
			}
		}
	}

	// Check general count-based rules
	if rule.MaxCount >= 0 && count > rule.MaxCount {
		return ruleResult{
			passed: false,
			message: fmt.Sprintf("Rule '%s' violated: Found %d matching vulnerabilities, maximum allowed is %d",
				rule.Name, count, rule.MaxCount),
		}
	}

	if rule.MinCount > 0 && count < rule.MinCount {
		return ruleResult{
			passed: false,
			message: fmt.Sprintf("Rule '%s' violated: Found %d matching vulnerabilities, minimum required is %d",
				rule.Name, count, rule.MinCount),
		}
	}

	return ruleResult{passed: true}
}

// matchesRule checks if a vulnerability matches a rule
func matchesRule(rule Rule, vuln Vulnerability) bool {
	// Check severity
	if rule.Severity != "" && !strings.EqualFold(vuln.Severity, rule.Severity) {
		return false
	}

	// Check package names
	if len(rule.Packages) > 0 {
		packageMatch := false
		for _, pkg := range rule.Packages {
			if strings.Contains(strings.ToLower(vuln.PackageName), strings.ToLower(pkg)) {
				packageMatch = true
				break
			}
		}
		if !packageMatch {
			return false
		}
	}

	// Check vulnerability IDs
	if len(rule.VulnIDs) > 0 {
		vulnMatch := false
		for _, vulnID := range rule.VulnIDs {
			if strings.Contains(strings.ToLower(vuln.VulnerabilityID), strings.ToLower(vulnID)) {
				vulnMatch = true
				break
			}
		}
		if !vulnMatch {
			return false
		}
	}

	return true
}

// countVulnerabilitiesBySeverity counts vulnerabilities by severity
func countVulnerabilitiesBySeverity(vulnerabilities []Vulnerability, severity string) int {
	count := 0
	for _, vuln := range vulnerabilities {
		if strings.EqualFold(vuln.Severity, severity) {
			count++
		}
	}
	return count
}

// CreateDefaultPolicy creates a default policy file
func CreateDefaultPolicy(filePath string) error {
	defaultPolicy := Policy{
		Rules: []Rule{
			{
				Name:        "Critical Vulnerabilities",
				Description: "Deny if any CRITICAL vulnerabilities are found",
				Severity:    "CRITICAL",
				MaxCount:    0,
				Action:      "deny",
			},
			{
				Name:        "High Vulnerabilities",
				Description: "Warn if more than 5 HIGH vulnerabilities are found",
				Severity:    "HIGH",
				MaxCount:    5,
				Action:      "warn",
			},
			{
				Name:        "Medium Vulnerabilities",
				Description: "Warn if more than 20 MEDIUM vulnerabilities are found",
				Severity:    "MEDIUM",
				MaxCount:    20,
				Action:      "warn",
			},
		},
	}

	var data []byte
	var err error

	if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		data, err = yaml.Marshal(defaultPolicy)
	} else {
		data, err = json.MarshalIndent(defaultPolicy, "", "  ")
	}

	if err != nil {
		return fmt.Errorf("failed to marshal default policy: %w", err)
	}

	return ioutil.WriteFile(filePath, data, 0644)
}
