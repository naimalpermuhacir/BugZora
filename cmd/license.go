package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"bugzora/pkg/report"
	"bugzora/pkg/vuln"
)

var licenseCmd = &cobra.Command{
	Use:   "license [target]",
	Short: "Scan for licenses in filesystem or repository",
	Long:  `Scans a filesystem path or repository for software licenses and license compliance issues.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]

		log.Printf("Scanning for licenses in: %s...", target)

		trivyArgs := buildTrivyArgs("fs", target)
		trivyArgs = append(trivyArgs, "--scanners", "license")

		scanReport, err := vuln.ScanFilesystemWithArgs(context.Background(), target, trivyArgs, quiet)
		if err != nil {
			log.Fatalf("License scan error: %v", err)
		}

		if err := report.WriteResults(scanReport.Results, outputFormat, target); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(licenseCmd)
}
