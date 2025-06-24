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
		output, _ := cmd.Flags().GetString("output")
		quiet, _ := cmd.Flags().GetBool("quiet")

		// Call the scanner function which wraps the trivy CLI.
		scanReport, err := vuln.ScanFilesystem(context.Background(), fsPath, quiet)
		if err != nil {
			log.Fatalf("Filesystem scan error: %v", err)
		}

		// Pass the results from the report to our reporting function.
		if err := report.WriteResults(scanReport.Results, output, fsPath); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fsCmd)
	fsCmd.Flags().StringP("output", "o", "table", "Output format (table, json, pdf)")
	fsCmd.Flags().BoolP("quiet", "q", false, "suppress progress messages")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
