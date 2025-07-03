package report

import (
	"github.com/aquasecurity/trivy/pkg/types"
)

// Demo modu - gerçek rapor oluşturulmaz
func GenerateTableReport(report types.Report, outputFormat string) (string, error) {
	return "Demo modu - gerçek rapor oluşturulmaz", nil
}

// Demo modu - gerçek rapor oluşturulmaz
func GenerateJSONReport(report types.Report) (string, error) {
	return "Demo modu - gerçek rapor oluşturulmaz", nil
}

// Demo modu - gerçek rapor oluşturulmaz
func GeneratePDFReport(report types.Report, outputPath string) error {
	return nil
}

// Demo modu - gerçek rapor oluşturulmaz
func GenerateSBOMReport(report types.Report, format string) (string, error) {
	return "Demo modu - gerçek rapor oluşturulmaz", nil
}
