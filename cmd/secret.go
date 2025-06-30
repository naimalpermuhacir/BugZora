package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"bugzora/pkg/report"
	"bugzora/pkg/vuln"
)

var secretCmd = &cobra.Command{
	Use:   "secret [target]",
	Short: "Scan for secrets in filesystem or repository",
	Long:  `Scans a filesystem path or repository for exposed secrets, API keys, passwords, and other sensitive information.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]

		log.Printf("Scanning for secrets in: %s...", target)

		trivyArgs := buildTrivyArgs("fs", target)
		trivyArgs = append(trivyArgs, "--scanners", "secret")

		scanReport, err := vuln.ScanFilesystemWithArgs(context.Background(), target, trivyArgs, quiet)
		if err != nil {
			log.Fatalf("Secret scan error: %v", err)
		}

		if err := report.WriteResults(scanReport.Results, outputFormat, target); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(secretCmd)
}
