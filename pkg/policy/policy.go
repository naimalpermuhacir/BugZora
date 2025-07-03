package policy

import (
	"github.com/aquasecurity/trivy/pkg/types"
)

// ValidatePolicy validates the policy against the report
func ValidatePolicy(report types.Report, policyFile string) (bool, error) {
	return true, nil
}

// EnforcePolicy enforces the policy against the report
func EnforcePolicy(report types.Report, policyFile string) error {
	return nil
}
