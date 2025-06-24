package report

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aquasecurity/trivy/pkg/types"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// PrintTable is the main entry point for printing vulnerability tables.
func PrintTable(target string, results types.Results) {
	if len(results) == 0 {
		fmt.Printf("✅ No vulnerabilities found for %s\n", target)
		return
	}
	fmt.Printf("\n--- Vulnerability Scan Report for: %s ---\n", target)

	for _, result := range results {
		if len(result.Vulnerabilities) == 0 {
			continue
		}
		printSingleResultTable(result)
	}
}

// printSingleResultTable renders a single, bordered table for a specific result set.
func printSingleResultTable(result types.Result) {
	header := fmt.Sprintf("🎯 Target: %s (%s)", result.Target, result.Type)
	fmt.Println(header)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Package", "Vulnerability ID", "Severity", "Installed Ver.", "Fixed Ver.", "Title", "Reference"})
	table.SetAutoWrapText(true) // Enable text wrapping for long titles/links
	table.SetRowLine(true)

	// Sort vulnerabilities by severity
	sort.Slice(result.Vulnerabilities, func(i, j int) bool {
		return severityToPriority(result.Vulnerabilities[i].Severity) > severityToPriority(result.Vulnerabilities[j].Severity)
	})

	severityCounts := make(map[string]int)
	for _, v := range result.Vulnerabilities {
		severityCounts[v.Severity]++

		// Get comprehensive reference links
		references := getMultipleReferences(v, result.Target, string(result.Type))

		row := []string{
			v.PkgName,
			v.VulnerabilityID,
			colorizeSeverity(v.Severity),
			v.InstalledVersion,
			v.FixedVersion,
			v.Title,
			references,
		}
		table.Append(row)
	}

	table.Render()
	fmt.Println(renderSummary(severityCounts))
}

// getMultipleReferences generates comprehensive reference links based on OS and vulnerability data
func getMultipleReferences(v types.DetectedVulnerability, target, osType string) string {
	var refs []string

	// Always include AquaSec if available
	if v.PrimaryURL != "" {
		refs = append(refs, fmt.Sprintf("🔍 AquaSec: %s", v.PrimaryURL))
	}

	// Add OS-specific references
	osRefs := getOSSpecificReferences(v, target, osType)
	refs = append(refs, osRefs...)

	// Add CVE database references
	cveRefs := getCVEReferences(v)
	refs = append(refs, cveRefs...)

	// Add NVD reference
	if strings.HasPrefix(v.VulnerabilityID, "CVE-") {
		refs = append(refs, fmt.Sprintf("📋 NVD: https://nvd.nist.gov/vuln/detail/%s", v.VulnerabilityID))
	}

	// If no references found, return N/A
	if len(refs) == 0 {
		return "N/A"
	}

	return strings.Join(refs, "\n")
}

// getOSSpecificReferences returns OS-specific reference links
func getOSSpecificReferences(v types.DetectedVulnerability, target, osType string) []string {
	var refs []string
	cveID := v.VulnerabilityID

	// Determine OS from target and type
	osName := strings.ToLower(osType)
	if strings.Contains(strings.ToLower(target), "ubuntu") {
		osName = "ubuntu"
	} else if strings.Contains(strings.ToLower(target), "debian") {
		osName = "debian"
	} else if strings.Contains(strings.ToLower(target), "alpine") {
		osName = "alpine"
	} else if strings.Contains(strings.ToLower(target), "centos") || strings.Contains(strings.ToLower(target), "rhel") {
		osName = "redhat"
	}

	// Add OS-specific references
	switch osName {
	case "ubuntu":
		refs = append(refs, fmt.Sprintf("🐧 Ubuntu: https://ubuntu.com/security/cve/%s", cveID))
		refs = append(refs, fmt.Sprintf("🔧 Ubuntu Tracker: https://ubuntu.com/security/%s", cveID))
	case "debian":
		refs = append(refs, fmt.Sprintf("📦 Debian: https://security-tracker.debian.org/tracker/%s", cveID))
		refs = append(refs, fmt.Sprintf("🔧 Debian Security: https://www.debian.org/security/%s", cveID))
	case "alpine":
		refs = append(refs, fmt.Sprintf("🏔️ Alpine: https://security.alpinelinux.org/vuln/%s", cveID))
	case "redhat":
		refs = append(refs, fmt.Sprintf("🔴 Red Hat: https://access.redhat.com/security/cve/%s", cveID))
		refs = append(refs, fmt.Sprintf("📋 Red Hat Bugzilla: https://bugzilla.redhat.com/show_bug.cgi?id=%s", cveID))
	}

	// Add existing vendor-specific references from Trivy data
	for _, ref := range v.References {
		if strings.Contains(ref, "debian.org/security") {
			refs = append(refs, fmt.Sprintf("📋 Debian Advisory: %s", ref))
		} else if strings.Contains(ref, "ubuntu.com/security/cve") {
			refs = append(refs, fmt.Sprintf("📋 Ubuntu Advisory: %s", ref))
		} else if strings.Contains(ref, "access.redhat.com/security/cve") {
			refs = append(refs, fmt.Sprintf("📋 Red Hat Advisory: %s", ref))
		} else if strings.Contains(ref, "alpinelinux.org") {
			refs = append(refs, fmt.Sprintf("📋 Alpine Advisory: %s", ref))
		}
	}

	return refs
}

// getCVEReferences returns CVE database references
func getCVEReferences(v types.DetectedVulnerability) []string {
	var refs []string
	cveID := v.VulnerabilityID

	if strings.HasPrefix(cveID, "CVE-") {
		refs = append(refs, fmt.Sprintf("📊 CVE Details: https://www.cvedetails.com/cve/%s/", cveID))
		refs = append(refs, fmt.Sprintf("🔍 MITRE: https://cve.mitre.org/cgi-bin/cvename.cgi?name=%s", cveID))
	}

	return refs
}

func renderSummary(counts map[string]int) string {
	total := 0
	var summaryParts []string

	severities := []string{"CRITICAL", "HIGH", "MEDIUM", "LOW", "UNKNOWN"}
	for _, sev := range severities {
		if count, ok := counts[sev]; ok && count > 0 {
			summaryParts = append(summaryParts, colorizeSeverityCount(sev, count))
			total += count
		}
	}

	return fmt.Sprintf("Summary: Total: %d (%s)", total, strings.Join(summaryParts, ", "))
}

func colorizeSeverityCount(severity string, count int) string {
	return fmt.Sprintf("%s: %d", colorizeSeverity(severity), count)
}

func colorizeSeverity(severity string) string {
	switch severity {
	case "CRITICAL":
		return color.New(color.FgRed, color.Bold).Sprint(severity)
	case "HIGH":
		return color.New(color.FgRed).Sprint(severity)
	case "MEDIUM":
		return color.New(color.FgYellow).Sprint(severity)
	case "LOW":
		return color.New(color.FgCyan).Sprint(severity)
	default:
		return color.New(color.FgWhite).Sprint(severity)
	}
}

func severityToPriority(severity string) int {
	switch severity {
	case "CRITICAL":
		return 5
	case "HIGH":
		return 4
	case "MEDIUM":
		return 3
	case "LOW":
		return 2
	default:
		return 1
	}
}
