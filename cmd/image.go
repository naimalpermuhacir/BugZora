/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"context"
	"io/fs"
	"log"

	"github.com/spf13/cobra"

	"bugzora/pkg/report" // We might need this for severity filtering later
	"bugzora/pkg/vuln"   // Even though trivy handles image scanning, we keep the same entrypoint
)

var imageCmd = &cobra.Command{
	Use:   "image [image-name]",
	Short: "Scan a container image for vulnerabilities",
	Long:  `Scans a given container image from a remote registry (like Docker Hub) for OS packages and their vulnerabilities.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]
		output, _ := cmd.Flags().GetString("output")
		quiet, _ := cmd.Flags().GetBool("quiet")

		// Call the scanner function which wraps the trivy CLI.
		log.Printf("Scanning image: %s... (this might take a while)", imageName)
		scanReport, err := vuln.ScanImage(context.Background(), imageName, quiet)
		if err != nil {
			log.Fatalf("Image scan error: %v", err)
		}

		// Pass the results from the report to our reporting function.
		if err := report.WriteResults(scanReport.Results, output, imageName); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.Flags().StringP("output", "o", "table", "Output format (table, json, pdf)")
	imageCmd.Flags().BoolP("quiet", "q", false, "suppress progress messages")
}

// LayerFS is a simple fs.FS implementation that overlays tar layers.
// NOTE: This is a simplified stand-in. The actual implementation is in go-containerregistry.
// We ensure we call the correct function: tarball.LayerFS(img).
type LayerFS struct {
	fs.FS
}
