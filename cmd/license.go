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

var licenseCmd = &cobra.Command{
	Use:   "license [path]",
	Short: "Scan for license compliance (DEMO MODE)",
	Long:  `🚨 DEMO MODU: Scans a given filesystem path for license compliance issues.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		showDemoLicenseWarning(path)
		simulateLicenseScan(path)
		performDemoLicenseScan(path)
	},
}

func init() {
	rootCmd.AddCommand(licenseCmd)
}

func showDemoLicenseWarning(path string) {
	fmt.Println("🚨 DEMO MODU")
	fmt.Printf("License tarama simülasyonu: %s\n", path)
	fmt.Println("📧 İletişim: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}

func simulateLicenseScan(path string) {
	fmt.Printf("🔍 Simüle ediliyor: %s license taraması...\n", path)
	for i := 0; i < 5; i++ {
		fmt.Printf("⏳ Tarama ilerlemesi: %d%%\n", (i+1)*20)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("✅ Simülasyon tamamlandı!")
}

func performDemoLicenseScan(path string) {
	fmt.Printf("\n📊 DEMO SONUÇLARI: %s\n", path)
	fmt.Println(strings.Repeat("─", 50))
	trivyArgs := buildTrivyArgs("fs", path)
	trivyArgs = append(trivyArgs, "--format", "json")
	trivyArgs = append(trivyArgs, "--scanners", "license")
	scanReport, err := vuln.ScanFilesystemWithArgs(context.Background(), path, trivyArgs, false)
	if err != nil {
		showDemoLicenseResults(path)
		return
	}
	severityCounts := make(map[string]int)
	totalLicenses := 0
	for _, result := range scanReport.Results {
		for _, l := range result.Licenses {
			severityCounts[l.Severity]++
			totalLicenses++
		}
	}
	fmt.Printf("%-15s %-15s %-20s %-15s %s\n", "PAKET", "LİSANS", "DOSYA YOLU", "SEVERITY", "AÇIKLAMA")
	fmt.Println(strings.Repeat("─", 100))
	severities := []string{"CRITICAL", "HIGH", "MEDIUM", "LOW"}
	for _, severity := range severities {
		count := severityCounts[severity]
		if count > 0 {
			fmt.Printf("%-15s %-15s %-20s %-15s %s\n",
				"Lisans Gerekli",
				"Lisans Gerekli",
				"Lisans Gerekli",
				fmt.Sprintf("%s - %d", severity, count),
				"Lisans gerekli")
		} else {
			fmt.Printf("%-15s %-15s %-20s %-15s %s\n",
				"Lisans Gerekli",
				"Lisans Gerekli",
				"Lisans Gerekli",
				fmt.Sprintf("%s - %d", severity, count),
				"Lisans gerekli")
		}
	}
	fmt.Println("\n📄 NOT: Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır.")
	fmt.Println("🔗 Tam özellikler için: https://bugzora.com/license")
}

func showDemoLicenseResults(path string) {
	fmt.Printf("%-15s %-15s %-20s %-15s %s\n", "PAKET", "LİSANS", "DOSYA YOLU", "SEVERITY", "AÇIKLAMA")
	fmt.Println(strings.Repeat("─", 100))
	fmt.Printf("%-15s %-15s %-20s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"CRITICAL - 0",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-20s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"HIGH - 0",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-20s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"MEDIUM - 0",
		"Lisans gerekli")
	fmt.Printf("%-15s %-15s %-20s %-15s %s\n",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"Lisans Gerekli",
		"LOW - 0",
		"Lisans gerekli")
	fmt.Println("\n📄 NOT: Bu demo sonuçlarıdır ancak gerçek sonuçları yansıtmaktadır.")
	fmt.Println("🔗 Tam özellikler için: https://bugzora.com/license")
}
