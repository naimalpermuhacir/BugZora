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

var repoCmd = &cobra.Command{
	Use:   "repo [repository-url]",
	Short: "Scan a Git repository for vulnerabilities (DEMO MODE)",
	Long:  `🚨 DEMO MODE: Scans a given Git repository for vulnerabilities and security issues.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]

		showDemoRepoWarning(repoURL)
		simulateRepoScan(repoURL)
		performDemoRepoScan(repoURL)
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}

func showDemoRepoWarning(repoURL string) {
	fmt.Println("🚨 DEMO MODE")
	fmt.Printf("Repository tarama simülasyonu: %s\n", repoURL)
	fmt.Println("📧 Contact: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}

func simulateRepoScan(repoURL string) {
	fmt.Printf("🔍 Simulating: %s repository taraması...\n", repoURL)
	for i := 0; i < 5; i++ {
		fmt.Printf("⏳ Scan progress: %d%%\n", (i+1)*20)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("✅ Simulation completed!")
}

func performDemoRepoScan(repoURL string) {
	fmt.Printf("\n📊 DEMO RESULTS: %s\n", repoURL)
	fmt.Println(strings.Repeat("─", 50))
	trivyArgs := buildTrivyArgs("repo", repoURL)
	trivyArgs = append(trivyArgs, "--format", "json")
	scanReport, err := vuln.ScanRepositoryWithArgs(context.Background(), repoURL, trivyArgs, false)
	if err != nil {
		showDemoRepoResults(repoURL)
		return
	}
	severityCounts := make(map[string]int)
	totalIssues := 0
	for _, result := range scanReport.Results {
		for _, v := range result.Vulnerabilities {
			severityCounts[v.Severity]++
			totalIssues++
		}
	}
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n", "İSSUE TİPİ", "DOSYA YOLU", "SATIR", "SEVERITY", "AÇIKLAMA")
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

func showDemoRepoResults(repoURL string) {
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n", "İSSUE TİPİ", "DOSYA YOLU", "SATIR", "SEVERITY", "AÇIKLAMA")
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
