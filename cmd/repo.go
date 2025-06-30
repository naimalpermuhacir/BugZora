package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"bugzora/pkg/report"
	"bugzora/pkg/vuln"
)

var repoCmd = &cobra.Command{
	Use:   "repo [repository-url]",
	Short: "Scan a Git repository for vulnerabilities, secrets, and licenses",
	Long:  `Scans a Git repository for vulnerabilities, secrets, licenses, and configuration issues.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]

		log.Printf("Scanning repository: %s...", repoURL)

		trivyArgs := buildTrivyArgs("repo", repoURL)

		scanReport, err := vuln.ScanRepositoryWithArgs(context.Background(), repoURL, trivyArgs, quiet)
		if err != nil {
			log.Fatalf("Repository scan error: %v", err)
		}

		if err := report.WriteResults(scanReport.Results, outputFormat, repoURL); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
 