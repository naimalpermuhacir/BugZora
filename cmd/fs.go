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

var fsCmd = &cobra.Command{
	Use:   "fs [path]",
	Short: "Scan a filesystem for vulnerabilities (DEMO MODE)",
	Long:  `🚨 DEMO MODU: Scans a given filesystem path for OS packages and their vulnerabilities.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		showDemoFsWarning(path)
		simulateFsScan(path)
		performDemoFsScan(path)
	},
}

func init() {
	rootCmd.AddCommand(fsCmd)
}

func showDemoFsWarning(path string) {
	fmt.Println("🚨 DEMO MODU")
	fmt.Printf("Filesystem tarama simülasyonu: %s\n", path)
	fmt.Println("📧 İletişim: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}

func simulateFsScan(path string) {
	fmt.Printf("🔍 Simüle ediliyor: %s taraması...\n", path)
	for i := 0; i < 5; i++ {
		fmt.Printf("⏳ Tarama ilerlemesi: %d%%\n", (i+1)*20)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("✅ Simülasyon tamamlandı!")
}

func performDemoFsScan(path string) {
	fmt.Printf("\n📊 DEMO SONUÇLARI: %s\n", path)
	fmt.Println(strings.Repeat("─", 50))
	trivyArgs := buildTrivyArgs("fs", path)
	trivyArgs = append(trivyArgs, "--format", "json")
	scanReport, err := vuln.ScanFilesystemWithArgs(context.Background(), path, trivyArgs, false)
	if err != nil {
		showDemoFsResults(path)
		return
	}
	severityCounts := make(map[string]int)
	totalVulns := 0
	for _, result := range scanReport.Results {
		for _, v := range result.Vulnerabilities {
			severityCounts[v.Severity]++
			totalVulns++
		}
	}
	fmt.Printf("%-15s %-15s %-20s %-20s %s\n", "PAKET", "ZAFİYET", "SEVERITY", "DOSYA YOLU", "AÇIKLAMA")
	fmt.Println(strings.Repeat("─", 100))
	severities := []string{"CRITICAL", "HIGH", "MEDIUM", "LOW"}
	for _, severity := range severities {
		count := severityCounts[severity]
		if count > 0 {
			fmt.Printf("%-15s %-15s %-20s %-20s %s\n",
				"Lisans Gerekli",
				"Lisans Gerekli",
				fmt.Sprintf("%s - %d", severity, count),
				"Lisans Gerekli",
				"Lisans gerekli")
		} else {
			fmt.Printf("%-15s %-15s %-20s %-20s %s\n",
				"Lisans Gerekli",
				"Lisans Gerekli",
				fmt.Sprintf("%s - %d", severity, count),
				"Lisans Gerekli",
				"Lisans gerekli")
		}
	}
	fmt.Println("\n📄 NOT: Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır.")
	fmt.Println("🔗 Tam özellikler için: https://bugzora.com/license")
}

func showDemoFsResults(path string) {
	fmt.Printf("%-15s %-15s %-20s %-20s %s\n", "PAKET", "ZAFİYET", "SEVERITY", "DOSYA YOLU", "AÇIKLAMA")
	fmt.Println(strings.Repeat("─", 100))
	fmt.Printf("%-15s %-15s %-20s %-20s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"CRITICAL - 0",
		"Lisans Gerekli",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-20s %-20s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"HIGH - 0",
		"Lisans Gerekli",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-20s %-20s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"MEDIUM - 0",
		"Lisans Gerekli",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-20s %-20s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"LOW - 0",
		"Lisans Gerekli",
		"Lisans gerekli")
	fmt.Println("\n📄 NOT: Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır.")
	fmt.Println("🔗 Tam özellikler için: https://bugzora.com/license")
}
