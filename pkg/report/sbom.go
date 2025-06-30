package report

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aquasecurity/trivy/pkg/types"
)

// CycloneDX SBOM structures
type CycloneDXBOM struct {
	BOMFormat       string                   `json:"bomFormat"`
	SpecVersion     string                   `json:"specVersion"`
	SerialNumber    string                   `json:"serialNumber"`
	Version         int                      `json:"version"`
	Metadata        CycloneDXMetadata        `json:"metadata"`
	Components      []CycloneDXComponent     `json:"components"`
	Vulnerabilities []CycloneDXVulnerability `json:"vulnerabilities,omitempty"`
}

type CycloneDXMetadata struct {
	Timestamp string             `json:"timestamp"`
	Tools     []CycloneDXTool    `json:"tools"`
	Component CycloneDXComponent `json:"component"`
}

type CycloneDXTool struct {
	Vendor  string `json:"vendor"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CycloneDXComponent struct {
	Type            string                   `json:"type"`
	BOMRef          string                   `json:"bom-ref"`
	Name            string                   `json:"name"`
	Version         string                   `json:"version"`
	PURL            string                   `json:"purl,omitempty"`
	CPE             string                   `json:"cpe,omitempty"`
	Licenses        []CycloneDXLicense       `json:"licenses,omitempty"`
	Vulnerabilities []CycloneDXVulnerability `json:"vulnerabilities,omitempty"`
}

type CycloneDXLicense struct {
	ID string `json:"id,omitempty"`
}

type CycloneDXVulnerability struct {
	ID          string               `json:"id"`
	Description string               `json:"description"`
	Ratings     []CycloneDXRating    `json:"ratings"`
	References  []CycloneDXReference `json:"references,omitempty"`
}

type CycloneDXRating struct {
	Severity string `json:"severity"`
	Method   string `json:"method"`
	Vector   string `json:"vector"`
}

type CycloneDXReference struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

// SPDX SBOM structures
type SPDXDocument struct {
	SPDXVersion       string            `json:"spdxVersion"`
	DataLicense       string            `json:"dataLicense"`
	SPDXID            string            `json:"spdxID"`
	DocumentName      string            `json:"documentName"`
	DocumentNamespace string            `json:"documentNamespace"`
	Creator           string            `json:"creator"`
	Created           string            `json:"created"`
	Packages          []SPDXPackage     `json:"packages"`
	ExternalRefs      []SPDXExternalRef `json:"externalRefs,omitempty"`
}

type SPDXPackage struct {
	SPDXID                  string            `json:"spdxID"`
	Name                    string            `json:"name"`
	VersionInfo             string            `json:"versionInfo"`
	PackageFileName         string            `json:"packageFileName"`
	PackageVerificationCode string            `json:"packageVerificationCode"`
	PackageLicenseConcluded string            `json:"packageLicenseConcluded"`
	PackageLicenseDeclared  string            `json:"packageLicenseDeclared"`
	PackageCopyrightText    string            `json:"packageCopyrightText"`
	ExternalRefs            []SPDXExternalRef `json:"externalRefs,omitempty"`
}

type SPDXExternalRef struct {
	ReferenceCategory string `json:"referenceCategory"`
	ReferenceType     string `json:"referenceType"`
	ReferenceLocator  string `json:"referenceLocator"`
}

// WriteCycloneDX generates a CycloneDX SBOM
func WriteCycloneDX(fileNameBase string, results types.Results) error {
	fileName := fmt.Sprintf("%s-cyclonedx.json", fileNameBase)

	bom := CycloneDXBOM{
		BOMFormat:    "CycloneDX",
		SpecVersion:  "1.5",
		SerialNumber: fmt.Sprintf("urn:uuid:%s", generateUUID()),
		Version:      1,
		Metadata: CycloneDXMetadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Tools: []CycloneDXTool{
				{
					Vendor:  "BugZora",
					Name:    "bugzora",
					Version: "1.3.0",
				},
			},
			Component: CycloneDXComponent{
				Type:    "application",
				BOMRef:  "root-component",
				Name:    "bugzora-scan",
				Version: "1.0.0",
			},
		},
		Components:      []CycloneDXComponent{},
		Vulnerabilities: []CycloneDXVulnerability{},
	}

	// Process scan results
	for _, result := range results {
		for _, vuln := range result.Vulnerabilities {
			// Create component
			component := CycloneDXComponent{
				Type:    "library",
				BOMRef:  fmt.Sprintf("pkg:%s/%s@%s", getPackageType(string(result.Type)), vuln.PkgName, vuln.InstalledVersion),
				Name:    vuln.PkgName,
				Version: vuln.InstalledVersion,
				PURL:    generatePURL(vuln.PkgName, vuln.InstalledVersion, string(result.Type)),
			}

			// Add vulnerability to component
			vulnerability := CycloneDXVulnerability{
				ID:          vuln.VulnerabilityID,
				Description: vuln.Title,
				Ratings: []CycloneDXRating{
					{
						Severity: convertSeverity(vuln.Severity),
						Method:   "other",
						Vector:   "CVSS:3.1",
					},
				},
			}

			// Add references
			if vuln.PrimaryURL != "" {
				vulnerability.References = append(vulnerability.References, CycloneDXReference{
					URL:  vuln.PrimaryURL,
					Type: "ADVISORY",
				})
			}

			component.Vulnerabilities = append(component.Vulnerabilities, vulnerability)
			bom.Components = append(bom.Components, component)
			bom.Vulnerabilities = append(bom.Vulnerabilities, vulnerability)
		}
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(bom, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal CycloneDX JSON: %w", err)
	}

	// Write to file
	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write CycloneDX file: %w", err)
	}

	fmt.Printf("✅ CycloneDX SBOM generated: %s\n", fileName)
	return nil
}

// WriteSPDX generates an SPDX SBOM (tag-value format, not JSON)
func WriteSPDX(fileNameBase string, results types.Results) error {
	fileName := fmt.Sprintf("%s-spdx.spdx", fileNameBase)

	// Simüle edilmiş bir SPDX tag-value çıktısı (örnek)
	spdxContent := `SPDXVersion: SPDX-2.3
DataLicense: CC0-1.0
SPDXID: SPDXRef-DOCUMENT
DocumentName: BugZora Security Scan SBOM
DocumentNamespace: https://bugzora.dev/sbom/` + generateUUID() + `
Creator: Tool: BugZora-1.3.0
Created: ` + time.Now().UTC().Format("2006-01-02T15:04:05Z") + `
`

	for _, result := range results {
		for _, vuln := range result.Vulnerabilities {
			spdxContent += fmt.Sprintf(`
##### Package: %s
PackageName: %s
SPDXID: SPDXRef-Package-%s-%s
PackageVersion: %s
`,
				vuln.PkgName, vuln.PkgName, strings.ReplaceAll(vuln.PkgName, "/", "-"), vuln.InstalledVersion, vuln.InstalledVersion)
		}
	}

	if err := os.WriteFile(fileName, []byte(spdxContent), 0644); err != nil {
		return fmt.Errorf("failed to write SPDX file: %w", err)
	}

	fmt.Printf("✅ SPDX SBOM (tag-value format) generated: %s\n", fileName)
	return nil
}

// Helper functions
func generateUUID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func getPackageType(resultType string) string {
	switch resultType {
	case "alpine":
		return "apk"
	case "debian":
		return "deb"
	case "ubuntu":
		return "deb"
	case "redhat":
		return "rpm"
	case "centos":
		return "rpm"
	case "amazon":
		return "rpm"
	case "oracle":
		return "rpm"
	case "photon":
		return "rpm"
	case "suse":
		return "rpm"
	case "cbl-mariner":
		return "rpm"
	case "wolfi":
		return "apk"
	case "chainguard":
		return "apk"
	default:
		return "generic"
	}
}

func generatePURL(pkgName, version string, resultType string) string {
	pkgType := getPackageType(resultType)

	// Clean package name for PURL
	cleanName := strings.ReplaceAll(pkgName, "_", "-")
	cleanName = strings.ReplaceAll(cleanName, " ", "-")

	return fmt.Sprintf("pkg:%s/%s@%s", pkgType, cleanName, version)
}

func convertSeverity(severity string) string {
	switch strings.ToUpper(severity) {
	case "CRITICAL":
		return "critical"
	case "HIGH":
		return "high"
	case "MEDIUM":
		return "medium"
	case "LOW":
		return "low"
	default:
		return "unknown"
	}
}

func generateVerificationCode(pkgName, version string) string {
	// Simple hash-like verification code
	return fmt.Sprintf("%x", len(pkgName+version))
}
