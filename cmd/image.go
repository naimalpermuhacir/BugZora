// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"context"
	"fmt"
	"io/fs"
	"log"

	"github.com/spf13/cobra"

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

		// Build Trivy command with all the flags
		trivyArgs := buildTrivyArgs("image", imageName)

		scanReport, err := vuln.ScanImageWithArgs(context.Background(), imageName, trivyArgs, quiet)
		if err != nil {
			log.Fatalf("Image scan error: %v", err)
		}

		if err := report.WriteResults(scanReport.Results, outputFormat, imageName); err != nil {
			log.Fatalf("Failed to write report: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	// Remove local flags since they're now global
}

// buildTrivyArgs builds the Trivy command arguments from global flags
func buildTrivyArgs(command, target string) []string {
	args := []string{command, target}

	// Add all the relevant flags
	if outputFormat != "" && outputFormat != "table" {
		args = append(args, "--format", outputFormat)
	}
	if quiet {
		args = append(args, "--quiet")
	}
	if severity != "" {
		args = append(args, "--severity", severity)
	}
	if exitCode != 0 {
		args = append(args, "--exit-code", fmt.Sprintf("%d", exitCode))
	}
	if ignoreUnfixed {
		args = append(args, "--ignore-unfixed")
	}
	if ignorePolicy != "" {
		args = append(args, "--ignore-policy", ignorePolicy)
	}
	if len(skipDirs) > 0 {
		for _, dir := range skipDirs {
			args = append(args, "--skip-dirs", dir)
		}
	}
	if len(skipFiles) > 0 {
		for _, file := range skipFiles {
			args = append(args, "--skip-files", file)
		}
	}
	if listAllPkgs {
		args = append(args, "--list-all-pkgs")
	}
	if offlineScan {
		args = append(args, "--offline-scan")
	}
	if scanners != "" {
		args = append(args, "--scanners", scanners)
	}
	if template != "" {
		args = append(args, "--template", template)
	}
	if configFile != "" {
		args = append(args, "--config", configFile)
	}
	if token != "" {
		args = append(args, "--token", token)
	}
	if proxy != "" {
		args = append(args, "--proxy", proxy)
	}
	if insecure {
		args = append(args, "--insecure")
	}
	if timeout != "" {
		args = append(args, "--timeout", timeout)
	}
	if downloadDBOnly {
		args = append(args, "--download-db-only")
	}
	if debug {
		args = append(args, "--debug")
	}
	if trace {
		args = append(args, "--trace")
	}
	if noProgress {
		args = append(args, "--no-progress")
	}
	if skipUpdate {
		args = append(args, "--skip-update")
	}
	if skipDBUpdate {
		args = append(args, "--skip-db-update")
	}
	if skipPolicyUpdate {
		args = append(args, "--skip-policy-update")
	}
	if securityChecks != "" && securityChecks != "vuln" {
		args = append(args, "--security-checks", securityChecks)
	}
	if compliance != "" {
		args = append(args, "--compliance", compliance)
	}
	if policy != "" {
		args = append(args, "--policy", policy)
	}
	if len(namespaces) > 0 {
		for _, ns := range namespaces {
			args = append(args, "--namespaces", ns)
		}
	}
	if format != "" {
		args = append(args, "--format", format)
	}
	if output != "" {
		args = append(args, "--output", output)
	}
	if len(severities) > 0 {
		for _, sev := range severities {
			args = append(args, "--severities", sev)
		}
	}
	if len(ignoreIDs) > 0 {
		for _, id := range ignoreIDs {
			args = append(args, "--ignore-ids", id)
		}
	}
	if ignoreFile != "" {
		args = append(args, "--ignore-file", ignoreFile)
	}
	if includeDevDeps {
		args = append(args, "--include-dev-deps")
	}
	if skipJavaDB {
		args = append(args, "--skip-java-db")
	}
	if skipUnfixed {
		args = append(args, "--skip-unfixed")
	}
	if onlyUpdate != "" {
		args = append(args, "--only-update", onlyUpdate)
	}
	if refresh {
		args = append(args, "--refresh")
	}
	if autoRefresh {
		args = append(args, "--auto-refresh")
	}
	if light {
		args = append(args, "--light")
	}

	return args
}

// LayerFS is a simple fs.FS implementation that overlays tar layers.
// NOTE: This is a simplified stand-in. The actual implementation is in go-containerregistry.
// We ensure we call the correct function: tarball.LayerFS(img).
type LayerFS struct {
	fs.FS
}
