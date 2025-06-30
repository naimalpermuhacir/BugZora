// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"context"
	"io/fs"
	"log"
	"os"

	"github.com/spf13/cobra"

	"bugzora/pkg/policy"
	"bugzora/pkg/report"
	"bugzora/pkg/vuln"
)

var imageCmd = &cobra.Command{
	Use:   "image [image-name]",
	Short: "Scan a container image for vulnerabilities",
	Long:  `Scans a given container image from a remote registry (like Docker Hub) for OS packages and their vulnerabilities.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]

		log.Printf("Scanning image: %s... (this might take a while)", imageName)

		trivyArgs := buildTrivyArgs("image", imageName)

		scanReport, err := vuln.ScanImageWithArgs(context.Background(), imageName, trivyArgs, quiet)
		if err != nil {
			log.Fatalf("Image scan error: %v", err)
		}

		if policyFile != "" {
			pol, err := policy.LoadPolicy(policyFile)
			if err != nil {
				log.Fatalf("Failed to load policy file: %v", err)
			}
			var vulns []policy.Vulnerability
			for _, result := range scanReport.Results {
				for _, v := range result.Vulnerabilities {
					vulns = append(vulns, policy.Vulnerability{
						VulnerabilityID: v.VulnerabilityID,
						Severity:        v.Severity,
						PackageName:     v.PkgName,
						PackageVersion:  v.InstalledVersion,
						Title:           v.Title,
						Description:     v.Description,
						Metadata:        map[string]string{"target": result.Target, "type": string(result.Type)},
					})
				}
			}
			res := policy.EvaluatePolicy(pol, vulns)
			if !res.Passed {
				log.Printf("\n\033[31mPolicy Violations Detected!\033[0m")
				for _, v := range res.Violations {
					log.Printf("- %s", v)
				}
				os.Exit(3)
			}
			if len(res.Warnings) > 0 {
				log.Printf("\n\033[33mPolicy Warnings:\033[0m")
				for _, w := range res.Warnings {
					log.Printf("- %s", w)
				}
			}
		}

		if err := report.WriteResults(scanReport.Results, outputFormat, imageName); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}

// LayerFS is a simple fs.FS implementation that overlays tar layers.
// NOTE: This is a simplified stand-in. The actual implementation is in go-containerregistry.
// We ensure we call the correct function: tarball.LayerFS(img).
type LayerFS struct {
	fs.FS
}
