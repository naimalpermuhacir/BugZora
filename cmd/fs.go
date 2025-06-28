// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"

	"bugzora/pkg/policy"
	"bugzora/pkg/report"
	"bugzora/pkg/vuln"
)

// fsCmd represents the fs command
var fsCmd = &cobra.Command{
	Use:   "fs [path]",
	Short: "Scan a filesystem for vulnerabilities",
	Long:  `Scans a given filesystem path for OS packages and their vulnerabilities.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fsPath := args[0]

		// Build Trivy command with all the flags
		trivyArgs := buildTrivyArgs("fs", fsPath)

		scanReport, err := vuln.ScanFilesystemWithArgs(context.Background(), fsPath, trivyArgs, quiet)
		if err != nil {
			log.Fatalf("Filesystem scan error: %v", err)
		}

		// Policy enforcement
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

		if err := report.WriteResults(scanReport.Results, outputFormat, fsPath); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fsCmd)
	// Remove local flags since they're now global
}
