// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bugzora",
	Short: "A powerful vulnerability scanner for container images and filesystems",
	Long: `Bugzora is a comprehensive security scanning tool that leverages the 
industry-standard Trivy engine to provide detailed vulnerability analysis.

Features:
• Container image scanning from Docker Hub and other registries
• Filesystem vulnerability scanning
• Multiple output formats (table, JSON, PDF)
• OS-specific vulnerability references
• Colored terminal output with detailed tables

Examples:
  bugzora image ubuntu:20.04
  bugzora fs /path/to/filesystem
  bugzora image alpine:latest --output json
  bugzora fs ./my-app --output pdf`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bugzora.yaml)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
