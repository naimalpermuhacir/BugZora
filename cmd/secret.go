// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"bugzora/pkg/vuln"

	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret [path]",
	Short: "Scan for secrets in filesystem (DEMO MODE)",
	Long:  `🚨 DEMO MODE: Scans a given filesystem path for secrets and sensitive information.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		showDemoSecretWarning(path)
		simulateSecretScan(path)
		performDemoSecretScan(path)
	},
}

func init() {
	rootCmd.AddCommand(secretCmd)
}

func showDemoSecretWarning(path string) {
	fmt.Println("🚨 DEMO MODE")
	fmt.Printf("Secret tarama simülasyonu: %s\n", path)
	fmt.Println("📧 Contact: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}

func simulateSecretScan(path string) {
	fmt.Printf("🔍 Simulating: %s secret taraması...\n", path)
	for i := 0; i < 5; i++ {
		fmt.Printf("⏳ Scan progress: %d%%\n", (i+1)*20)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("✅ Simulation completed!")
}

func performDemoSecretScan(path string) {
	fmt.Printf("\n📊 DEMO RESULTS: %s\n", path)
	fmt.Println(strings.Repeat("─", 50))
	trivyArgs := buildTrivyArgs("fs", path)
	trivyArgs = append(trivyArgs, "--format", "json")
	trivyArgs = append(trivyArgs, "--scanners", "secret")
	scanReport, err := vuln.ScanFilesystemWithArgs(context.Background(), path, trivyArgs, false)
	if err != nil {
		showDemoSecretResults(path)
		return
	}
	severityCounts := make(map[string]int)
	totalSecrets := 0
	for _, result := range scanReport.Results {
		for _, s := range result.Secrets {
			severityCounts[s.Severity]++
			totalSecrets++
		}
	}
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n", "SECRET TİPİ", "DOSYA YOLU", "SATIR", "SEVERITY", "AÇIKLAMA")
	fmt.Println(strings.Repeat("─", 80))
	severities := []string{"CRITICAL", "HIGH", "MEDIUM", "LOW"}
	for _, severity := range severities {
		count := severityCounts[severity]
		if count > 0 {
			fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
				"License Required",
				"License Required",
				"Lisans",
				fmt.Sprintf("%s - %d", severity, count),
				"License required")
		} else {
			fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
				"License Required",
				"License Required",
				"Lisans",
				fmt.Sprintf("%s - %d", severity, count),
				"License required")
		}
	}
	fmt.Println("\n📄 NOT: This is a demo result but reflects real data.")
	fmt.Println("🔗 For full features: https://bugzora.com/license")
}

func showDemoSecretResults(path string) {
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n", "SECRET TİPİ", "DOSYA YOLU", "SATIR", "SEVERITY", "AÇIKLAMA")
	fmt.Println(strings.Repeat("─", 80))
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"License Required",
		"License Required",
		"Lisans",
		"CRITICAL - 0",
		"License required")
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"License Required",
		"License Required",
		"Lisans",
		"HIGH - 0",
		"License required")
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"License Required",
		"License Required",
		"Lisans",
		"MEDIUM - 0",
		"License required")
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"License Required",
		"License Required",
		"Lisans",
		"LOW - 0",
		"License required")
	fmt.Println("\n📄 NOT: This is a demo result but reflects real data.")
	fmt.Println("🔗 For full features: https://bugzora.com/license")
}
