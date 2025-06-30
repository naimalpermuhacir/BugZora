package report

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/trivy/pkg/types"
	"github.com/jung-kurt/gofpdf"
)

// WriteResults takes the scan results and outputs them in the specified format.
func WriteResults(results types.Results, format string, outputTarget string) error {
	sanitizedTarget := strings.ReplaceAll(outputTarget, "/", "_")
	sanitizedTarget = strings.ReplaceAll(sanitizedTarget, ":", "-")
	outputFileNameBase := "report-" + sanitizedTarget

	switch format {
	case "table":
		PrintTable(outputTarget, results)
		return nil
	case "json":
		return WriteJSON(outputFileNameBase, results)
	case "pdf":
		return WritePDF(outputFileNameBase, results)
	case "cyclonedx":
		return WriteCycloneDX(outputFileNameBase, results)
	case "spdx":
		return WriteSPDX(outputFileNameBase, results)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

// VulnerabilityReport represents the complete scan report structure for JSON
type VulnerabilityReport struct {
	ScanInfo   ScanInfo              `json:"scan_info"`
	Target     string                `json:"target"`
	ScanTime   time.Time             `json:"scan_time"`
	Summary    Summary               `json:"summary"`
	Results    []VulnerabilityResult `json:"results"`
	TotalCount int                   `json:"total_count"`
}

// ScanInfo contains metadata about the scan
type ScanInfo struct {
	Scanner  string    `json:"scanner"`
	Version  string    `json:"version"`
	ScanTime time.Time `json:"scan_time"`
	Target   string    `json:"target"`
}

// Summary contains vulnerability statistics
type Summary struct {
	Critical int `json:"critical"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`
	Unknown  int `json:"unknown"`
	Total    int `json:"total"`
}

// VulnerabilityResult represents a single scan result
type VulnerabilityResult struct {
	Target          string                `json:"target"`
	Type            string                `json:"type"`
	Vulnerabilities []VulnerabilityDetail `json:"vulnerabilities"`
	Summary         Summary               `json:"summary"`
}

// VulnerabilityDetail represents detailed vulnerability information
type VulnerabilityDetail struct {
	VulnerabilityID  string      `json:"vulnerability_id"`
	PkgName          string      `json:"package_name"`
	InstalledVersion string      `json:"installed_version"`
	FixedVersion     string      `json:"fixed_version"`
	Severity         string      `json:"severity"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	References       []Reference `json:"references"`
	PrimaryURL       string      `json:"primary_url"`
}

// Reference represents a single reference link
type Reference struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// WriteJSON generates a comprehensive JSON report
func WriteJSON(fileNameBase string, results types.Results) error {
	fileName := fmt.Sprintf("%s-json.json", fileNameBase)

	// Create the report structure
	report := VulnerabilityReport{
		ScanInfo: ScanInfo{
			Scanner:  "bugzora",
			Version:  "1.0.0",
			ScanTime: time.Now(),
		},
		ScanTime: time.Now(),
		Results:  []VulnerabilityResult{},
	}

	// Process each result
	for _, result := range results {
		if len(result.Vulnerabilities) == 0 {
			continue
		}

		vulnResult := VulnerabilityResult{
			Target:          result.Target,
			Type:            string(result.Type),
			Vulnerabilities: []VulnerabilityDetail{},
		}

		// Count vulnerabilities by severity
		severityCounts := make(map[string]int)

		for _, vuln := range result.Vulnerabilities {
			severityCounts[vuln.Severity]++

			// Generate comprehensive references
			references := generateReferences(vuln, result.Target, string(result.Type))

			vulnDetail := VulnerabilityDetail{
				VulnerabilityID:  vuln.VulnerabilityID,
				PkgName:          vuln.PkgName,
				InstalledVersion: vuln.InstalledVersion,
				FixedVersion:     vuln.FixedVersion,
				Severity:         vuln.Severity,
				Title:            vuln.Title,
				PrimaryURL:       vuln.PrimaryURL,
				References:       references,
			}

			vulnResult.Vulnerabilities = append(vulnResult.Vulnerabilities, vulnDetail)
		}

		// Calculate summary for this result
		vulnResult.Summary = Summary{
			Critical: severityCounts["CRITICAL"],
			High:     severityCounts["HIGH"],
			Medium:   severityCounts["MEDIUM"],
			Low:      severityCounts["LOW"],
			Unknown:  severityCounts["UNKNOWN"],
			Total:    len(result.Vulnerabilities),
		}

		report.Results = append(report.Results, vulnResult)
	}

	// Calculate overall summary
	overallSummary := Summary{}
	for _, result := range report.Results {
		overallSummary.Critical += result.Summary.Critical
		overallSummary.High += result.Summary.High
		overallSummary.Medium += result.Summary.Medium
		overallSummary.Low += result.Summary.Low
		overallSummary.Unknown += result.Summary.Unknown
		overallSummary.Total += result.Summary.Total
	}

	report.Summary = overallSummary
	report.TotalCount = overallSummary.Total

	// Marshal to JSON with pretty formatting
	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to file
	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	fmt.Printf("✅ JSON report generated: %s\n", fileName)
	return nil
}

// WritePDF generates a comprehensive PDF report
func WritePDF(fileNameBase string, results types.Results) error {
	fileName := fmt.Sprintf("%s-pdf.pdf", fileNameBase)

	// Create PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "B", 16)

	// Add title
	pdf.Cell(190, 10, "Security Scan Report")
	pdf.Ln(15)

	// Add scan information
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 8, "Scan Information")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(50, 6, "Scan Time:")
	pdf.Cell(140, 6, time.Now().Format("2006-01-02 15:04:05"))
	pdf.Ln(8)

	pdf.Cell(50, 6, "Scanner:")
	pdf.Cell(140, 6, "bugzora v1.0.0")
	pdf.Ln(8)

	// Calculate summary statistics
	summary := calculatePDFSummary(results)

	// Add summary section
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 8, "Summary Statistics")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(50, 6, "Critical:")
	pdf.Cell(20, 6, strconv.Itoa(summary.Critical))
	pdf.Ln(8)

	pdf.Cell(50, 6, "High:")
	pdf.Cell(20, 6, strconv.Itoa(summary.High))
	pdf.Ln(8)

	pdf.Cell(50, 6, "Medium:")
	pdf.Cell(20, 6, strconv.Itoa(summary.Medium))
	pdf.Ln(8)

	pdf.Cell(50, 6, "Low:")
	pdf.Cell(20, 6, strconv.Itoa(summary.Low))
	pdf.Ln(8)

	pdf.Cell(50, 6, "Unknown:")
	pdf.Cell(20, 6, strconv.Itoa(summary.Unknown))
	pdf.Ln(8)

	pdf.Cell(50, 6, "Total:")
	pdf.Cell(20, 6, strconv.Itoa(summary.Total))
	pdf.Ln(15)

	// Process each result
	for _, result := range results {
		if len(result.Vulnerabilities) == 0 {
			continue
		}

		// Add result section header
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(190, 8, fmt.Sprintf("Target: %s (%s)", result.Target, result.Type))
		pdf.Ln(10)

		// Create vulnerability table
		createVulnerabilityTable(pdf, result.Vulnerabilities)
		pdf.Ln(10)
	}

	// Save PDF
	if err := pdf.OutputFileAndClose(fileName); err != nil {
		return fmt.Errorf("failed to write PDF file: %w", err)
	}

	fmt.Printf("✅ PDF report generated: %s\n", fileName)
	return nil
}

// calculatePDFSummary calculates overall vulnerability statistics
func calculatePDFSummary(results types.Results) Summary {
	summary := Summary{}

	for _, result := range results {
		for _, vuln := range result.Vulnerabilities {
			switch vuln.Severity {
			case "CRITICAL":
				summary.Critical++
			case "HIGH":
				summary.High++
			case "MEDIUM":
				summary.Medium++
			case "LOW":
				summary.Low++
			default:
				summary.Unknown++
			}
			summary.Total++
		}
	}

	return summary
}

// createVulnerabilityTable creates a table of vulnerabilities in the PDF
func createVulnerabilityTable(pdf *gofpdf.Fpdf, vulnerabilities []types.DetectedVulnerability) {
	widths := []float64{30, 35, 25, 25, 20, 55}

	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(200, 200, 200)

	pdf.CellFormat(widths[0], 6, "Vuln ID", "1", 0, "", true, 0, "")
	pdf.CellFormat(widths[1], 6, "Package", "1", 0, "", true, 0, "")
	pdf.CellFormat(widths[2], 6, "Installed", "1", 0, "", true, 0, "")
	pdf.CellFormat(widths[3], 6, "Fixed", "1", 0, "", true, 0, "")
	pdf.CellFormat(widths[4], 6, "Severity", "1", 0, "", true, 0, "")
	pdf.CellFormat(widths[5], 6, "Title", "1", 0, "", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 8)

	for _, vuln := range vulnerabilities {
		fillColor := getSeverityColor(vuln.Severity)
		pdf.SetFillColor(fillColor[0], fillColor[1], fillColor[2])

		title := truncateText(vuln.Title, 40)
		pkgName := truncateText(vuln.PkgName, 25)
		installedVer := truncateText(vuln.InstalledVersion, 20)
		fixedVer := truncateText(vuln.FixedVersion, 20)

		pdf.CellFormat(widths[0], 6, vuln.VulnerabilityID, "1", 0, "", true, 0, "")
		pdf.CellFormat(widths[1], 6, pkgName, "1", 0, "", true, 0, "")
		pdf.CellFormat(widths[2], 6, installedVer, "1", 0, "", true, 0, "")
		pdf.CellFormat(widths[3], 6, fixedVer, "1", 0, "", true, 0, "")
		pdf.CellFormat(widths[4], 6, vuln.Severity, "1", 0, "", true, 0, "")
		pdf.CellFormat(widths[5], 6, title, "1", 0, "", true, 0, "")
		pdf.Ln(-1)

		if vuln.PrimaryURL != "" {
			pdf.SetFont("Arial", "I", 7)
			pdf.SetTextColor(0, 0, 255)
			refText := fmt.Sprintf("Reference: %s", truncateText(vuln.PrimaryURL, 80))
			pdf.Cell(190, 4, refText)
			pdf.Ln(5)
			pdf.SetTextColor(0, 0, 0)
			pdf.SetFont("Arial", "", 8)
		}
	}
}

// getSeverityColor returns RGB color values based on severity
func getSeverityColor(severity string) [3]int {
	switch severity {
	case "CRITICAL":
		return [3]int{255, 0, 0}
	case "HIGH":
		return [3]int{255, 165, 0}
	case "MEDIUM":
		return [3]int{255, 255, 0}
	case "LOW":
		return [3]int{0, 255, 0}
	default:
		return [3]int{200, 200, 200}
	}
}

// truncateText truncates text to specified length
func truncateText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	return text[:maxLength-3] + "..."
}

// generateReferences creates comprehensive reference links for a vulnerability
func generateReferences(vuln types.DetectedVulnerability, target, osType string) []Reference {
	var refs []Reference

	// Add AquaSec reference
	if vuln.PrimaryURL != "" {
		refs = append(refs, Reference{
			Type:        "Primary",
			URL:         vuln.PrimaryURL,
			Description: "Primary vulnerability analysis and details",
		})
	}

	// Add OS-specific references
	osRefs := generateOSReferences(vuln, target, osType)
	refs = append(refs, osRefs...)

	// Add CVE database references
	cveRefs := generateCVEReferences(vuln)
	refs = append(refs, cveRefs...)

	return refs
}

// generateOSReferences creates OS-specific reference links
func generateOSReferences(vuln types.DetectedVulnerability, target, osType string) []Reference {
	var refs []Reference
	cveID := vuln.VulnerabilityID

	// Determine OS from target and type
	osName := osType
	if osName == "" {
		osName = "unknown"
	}

	// Add OS-specific references
	switch osName {
	case "ubuntu":
		refs = append(refs, Reference{
			Type:        "Ubuntu Advisory",
			URL:         fmt.Sprintf("https://ubuntu.com/security/cve/%s", cveID),
			Description: "Ubuntu security advisory for " + cveID,
		})
	case "debian":
		refs = append(refs, Reference{
			Type:        "Debian Advisory",
			URL:         fmt.Sprintf("https://security-tracker.debian.org/tracker/%s", cveID),
			Description: "Debian security tracker for " + cveID,
		})
	case "alpine":
		refs = append(refs, Reference{
			Type:        "Alpine Advisory",
			URL:         fmt.Sprintf("https://alpinelinux.org/security/cve/%s", cveID),
			Description: "Alpine security advisory for " + cveID,
		})
	case "redhat":
		refs = append(refs, Reference{
			Type:        "Red Hat Advisory",
			URL:         fmt.Sprintf("https://access.redhat.com/security/cve/%s", cveID),
			Description: "Red Hat security advisory for " + cveID,
		})
	}

	return refs
}

// generateCVEReferences creates CVE database reference links
func generateCVEReferences(vuln types.DetectedVulnerability) []Reference {
	var refs []Reference
	cveID := vuln.VulnerabilityID

	if strings.HasPrefix(cveID, "CVE-") {
		refs = append(refs, Reference{
			Type:        "CVE Details",
			URL:         fmt.Sprintf("https://www.cvedetails.com/cve/%s/", cveID),
			Description: "Detailed CVE information and analysis",
		})
		refs = append(refs, Reference{
			Type:        "MITRE",
			URL:         fmt.Sprintf("https://cve.mitre.org/cgi-bin/cvename.cgi?name=%s", cveID),
			Description: "MITRE CVE database entry",
		})
	}

	return refs
}
