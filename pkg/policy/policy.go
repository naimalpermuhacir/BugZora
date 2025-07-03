package policy

import (
	"github.com/aquasecurity/trivy/pkg/types"
)

// Demo modu - gerçek policy kontrolü yapılmaz
func ValidatePolicy(report types.Report, policyFile string) (bool, error) {
	return true, nil
}

// Demo modu - gerçek policy kontrolü yapılmaz
func EnforcePolicy(report types.Report, policyFile string) error {
	return nil
}
