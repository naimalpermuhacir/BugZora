// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version information - set by GoReleaser
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

// Global flags that will be available to all commands
var (
	outputFormat     string
	quiet            bool
	severity         string
	exitCode         int
	ignoreUnfixed    bool
	ignorePolicy     string
	skipDirs         []string
	skipFiles        []string
	listAllPkgs      bool
	offlineScan      bool
	scanners         string
	template         string
	configFile       string
	token            string
	proxy            string
	insecure         bool
	timeout          string
	downloadDBOnly   bool
	debug            bool
	trace            bool
	noProgress       bool
	skipUpdate       bool
	skipDBUpdate     bool
	skipPolicyUpdate bool
	securityChecks   string
	compliance       string
	trivyPolicy      string
	namespaces       []string
	format           string
	output           string
	severities       []string
	ignoreIDs        []string
	ignoreFile       string
	includeDevDeps   bool
	skipJavaDB       bool
	skipUnfixed      bool
	onlyUpdate       string
	refresh          bool
	autoRefresh      bool
	light            bool
	policyFile       string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bugzora",
	Short: "BugZora - Container and filesystem security scanner (DEMO MODE)",
	Long: `🚨 DEMO MODU

BugZora is a comprehensive security scanner for container images and filesystems.
It uses Trivy as its scanning engine and provides enhanced reporting capabilities.

Features:
- Container image vulnerability scanning (DEMO)
- Filesystem security analysis (DEMO)
- Multiple output formats (table, JSON, PDF, SBOM) (DEMO)
- Policy enforcement (DEMO)
- Comprehensive vulnerability references (DEMO)
- Docker integration (DEMO)

Examples:
  bugzora image alpine:latest
  bugzora fs /path/to/filesystem
  bugzora image nginx:latest --format json --output report.json`,
	Version: fmt.Sprintf("%s (commit: %s, date: %s) - DEMO MODE", version, commit, date),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Persistent flags that will be available to all commands
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, json, pdf, sarif, cyclonedx, spdx)")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppress progress messages")
	rootCmd.PersistentFlags().StringVar(&severity, "severity", "", "Severities of security issues to be displayed (comma separated)")
	rootCmd.PersistentFlags().IntVar(&exitCode, "exit-code", 0, "Exit code when vulnerabilities were found")
	rootCmd.PersistentFlags().BoolVar(&ignoreUnfixed, "ignore-unfixed", false, "Display only fixed vulnerabilities")
	rootCmd.PersistentFlags().StringVar(&ignorePolicy, "ignore-policy", "", "Specify the Rego file path to evaluate each vulnerability")
	rootCmd.PersistentFlags().StringSliceVar(&skipDirs, "skip-dirs", []string{}, "Specify the directories where the traversal is skipped")
	rootCmd.PersistentFlags().StringSliceVar(&skipFiles, "skip-files", []string{}, "Specify the file paths to skip")
	rootCmd.PersistentFlags().BoolVar(&listAllPkgs, "list-all-pkgs", false, "Enabling the option will output all packages regardless of vulnerability")
	rootCmd.PersistentFlags().BoolVar(&offlineScan, "offline-scan", false, "Do not issue API requests to identify dependencies")
	rootCmd.PersistentFlags().StringVar(&scanners, "scanners", "", "Comma-separated list of what security issues to detect (vuln,config,secret,license)")
	rootCmd.PersistentFlags().StringVar(&template, "template", "", "Output template")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Config path")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "For authentication in client/server mode")
	rootCmd.PersistentFlags().StringVar(&proxy, "proxy", "", "HTTP proxy URL")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "Allow insecure server connections when using TLS")
	rootCmd.PersistentFlags().StringVar(&timeout, "timeout", "", "Timeout (e.g. 5s, 2m, 1h)")
	rootCmd.PersistentFlags().BoolVar(&downloadDBOnly, "download-db-only", false, "Download/update vulnerability database but don't run a scan")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolVar(&trace, "trace", false, "Enable trace output")
	rootCmd.PersistentFlags().BoolVar(&noProgress, "no-progress", false, "Suppress progress bar")
	rootCmd.PersistentFlags().BoolVar(&skipUpdate, "skip-update", false, "Skip db update")
	rootCmd.PersistentFlags().BoolVar(&skipDBUpdate, "skip-db-update", false, "Skip updating vulnerability database")
	rootCmd.PersistentFlags().BoolVar(&skipPolicyUpdate, "skip-policy-update", false, "Skip updating policy database")
	rootCmd.PersistentFlags().StringVar(&securityChecks, "security-checks", "vuln", "Comma-separated list of what security issues to detect (vuln,config,secret,license)")
	rootCmd.PersistentFlags().StringVar(&compliance, "compliance", "", "Compliance report to generate (docker-cis,kubernetes-cis,aws-cis)")
	rootCmd.PersistentFlags().StringVar(&trivyPolicy, "policy", "", "Specify the Rego file path to evaluate each vulnerability")
	rootCmd.PersistentFlags().StringSliceVar(&namespaces, "namespaces", []string{}, "Rego namespaces")
	rootCmd.PersistentFlags().StringVar(&format, "format", "", "Format (table, json, sarif, template, cyclonedx, spdx)")
	rootCmd.PersistentFlags().StringVar(&output, "output-file", "", "Output file path")
	rootCmd.PersistentFlags().StringSliceVar(&severities, "severities", []string{}, "Severities of security issues to be displayed (comma separated)")
	rootCmd.PersistentFlags().StringSliceVar(&ignoreIDs, "ignore-ids", []string{}, "Vulnerability IDs to ignore")
	rootCmd.PersistentFlags().StringVar(&ignoreFile, "ignore-file", "", "Specify .trivyignore file")
	rootCmd.PersistentFlags().BoolVar(&includeDevDeps, "include-dev-deps", false, "Include development dependencies in scanning")
	rootCmd.PersistentFlags().BoolVar(&skipJavaDB, "skip-java-db", false, "Skip updating Java index database")
	rootCmd.PersistentFlags().BoolVar(&skipUnfixed, "skip-unfixed", false, "Skip unfixed vulnerabilities")
	rootCmd.PersistentFlags().StringVar(&onlyUpdate, "only-update", "", "Update only specified database (e.g. alpine, debian, ubuntu)")
	rootCmd.PersistentFlags().BoolVar(&refresh, "refresh", false, "Refresh OVAL database")
	rootCmd.PersistentFlags().BoolVar(&autoRefresh, "auto-refresh", false, "Auto refresh before scan")
	rootCmd.PersistentFlags().BoolVar(&light, "light", false, "Light mode (disable security checks)")
	rootCmd.PersistentFlags().StringVar(&policyFile, "policy-file", "", "Path to policy file for enforcement")

	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(fsCmd)
	rootCmd.AddCommand(secretCmd)
	rootCmd.AddCommand(licenseCmd)
	rootCmd.AddCommand(repoCmd)
}
