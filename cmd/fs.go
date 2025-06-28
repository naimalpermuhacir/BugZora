// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

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

		if err := report.WriteResults(scanReport.Results, outputFormat, fsPath); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fsCmd)
	// Remove local flags since they're now global
}
