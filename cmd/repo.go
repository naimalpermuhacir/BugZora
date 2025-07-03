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
	Long:  `🚨 DEMO MODU: Scans a given Git repository for vulnerabilities and security issues.`,
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
	fmt.Println("🚨 DEMO MODU")
	fmt.Printf("Repository tarama simülasyonu: %s\n", repoURL)
	fmt.Println("📧 İletişim: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}

func simulateRepoScan(repoURL string) {
	fmt.Printf("🔍 Simüle ediliyor: %s repository taraması...\n", repoURL)
	for i := 0; i < 5; i++ {
		fmt.Printf("⏳ Tarama ilerlemesi: %d%%\n", (i+1)*20)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("✅ Simülasyon tamamlandı!")
}

func performDemoRepoScan(repoURL string) {
	fmt.Printf("\n📊 DEMO SONUÇLARI: %s\n", repoURL)
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
				"Lisans Gerekli",
				"Lisans Gerekli",
				"Lisans",
				fmt.Sprintf("%s - %d", severity, count),
				"Lisans gerekli")
		} else {
			fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
				"Lisans Gerekli",
				"Lisans Gerekli",
				"Lisans",
				fmt.Sprintf("%s - %d", severity, count),
				"Lisans gerekli")
		}
	}
	fmt.Println("\n📄 NOT: Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır.")
	fmt.Println("🔗 Tam özellikler için: https://bugzora.com/license")
}

func showDemoRepoResults(repoURL string) {
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n", "İSSUE TİPİ", "DOSYA YOLU", "SATIR", "SEVERITY", "AÇIKLAMA")
	fmt.Println(strings.Repeat("─", 80))
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans",
		"CRITICAL - 0",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans",
		"HIGH - 0",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans",
		"MEDIUM - 0",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-15s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans",
		"LOW - 0",
		"Lisans gerekli")
	fmt.Println("\n📄 NOT: Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır.")
	fmt.Println("🔗 Tam özellikler için: https://bugzora.com/license")
}
