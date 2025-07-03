package report

import (
	"github.com/aquasecurity/trivy/pkg/types"
)

// GenerateTableReport generates a table report from the scan results
func GenerateTableReport(report types.Report, outputFormat string) (string, error) {
	return "Demo mode - real report not generated", nil
}

// GenerateJSONReport generates a JSON report from the scan results
func GenerateJSONReport(report types.Report) (string, error) {
	return "Demo mode - real report not generated", nil
}

// GeneratePDFReport generates a PDF report from the scan results
func GeneratePDFReport(report types.Report, outputPath string) error {
	return nil
}

// GenerateSBOMReport generates an SBOM report from the scan results
func GenerateSBOMReport(report types.Report, format string) (string, error) {
	return "Demo mode - real report not generated", nil
}
