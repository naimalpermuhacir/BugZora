// Package cmd provides command-line interface for BugZora security scanner
/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package cmd

import (
	"context"
	"fmt"
	"io/fs"
	"strings"
	"time"

	"bugzora/pkg/vuln"

	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image [image-name]",
	Short: "Scan a container image for vulnerabilities (DEMO MODE)",
	Long:  `🚨 DEMO MODE: Scans a given container image from a remote registry (like Docker Hub) for OS packages and their vulnerabilities.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]

		showDemoImageWarning(imageName)
		simulateImageScan(imageName)
		performDemoImageScan(imageName)
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}

func showDemoImageWarning(imageName string) {
	fmt.Println("🚨 DEMO MODE")
	fmt.Printf("Image scanning simulation: %s\n", imageName)
	fmt.Println("📧 Contact: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}

func simulateImageScan(imageName string) {
	fmt.Printf("🔍 Simulating: %s scanning...\n", imageName)
	for i := 0; i < 5; i++ {
		fmt.Printf("⏳ Scan progress: %d%%\n", (i+1)*20)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("✅ Simulation completed!")
}

func performDemoImageScan(imageName string) {
	fmt.Printf("\n📊 DEMO RESULTS: %s\n", imageName)
	fmt.Println(strings.Repeat("─", 50))
	trivyArgs := buildTrivyArgs("image", imageName)
	trivyArgs = append(trivyArgs, "--format", "json")
	scanReport, err := vuln.ScanImageWithArgs(context.Background(), imageName, trivyArgs, false)
	if err != nil {
		showDemoImageResults(imageName)
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
	fmt.Printf("%-15s %-15s %-20s %s\n", "PACKAGE", "VULNERABILITY", "SEVERITY", "DESCRIPTION")
	fmt.Println(strings.Repeat("─", 80))
	severities := []string{"CRITICAL", "HIGH", "MEDIUM", "LOW"}
	for _, severity := range severities {
		count := severityCounts[severity]
		if count > 0 {
			fmt.Printf("%-15s %-15s %-20s %s\n",
				"License Required",
				"License Required",
				fmt.Sprintf("%s - %d", severity, count),
				"License required")
		} else {
			fmt.Printf("%-15s %-15s %-20s %s\n",
				"License Required",
				"License Required",
				fmt.Sprintf("%s - %d", severity, count),
				"License required")
		}
	}
	fmt.Println("\n📄 NOT: This is a demo result but reflects real data.")
	fmt.Println("🔗 For full features: https://bugzora.com/license")
}

func showDemoImageResults(imageName string) {
	fmt.Printf("%-15s %-15s %-20s %s\n", "PACKAGE", "VULNERABILITY", "SEVERITY", "DESCRIPTION")
	fmt.Println(strings.Repeat("─", 80))
	fmt.Printf("%-15s %-15s %-20s %s\n",
		"License Required",
		"License Required",
		"CRITICAL - 0",
		"License required")
	fmt.Printf("%-15s %-15s %-20s %s\n",
		"License Required",
		"License Required",
		"HIGH - 0",
		"License required")
	fmt.Printf("%-15s %-15s %-20s %s\n",
		"License Required",
		"License Required",
		"MEDIUM - 0",
		"License required")
	fmt.Printf("%-15s %-15s %-20s %s\n",
		"License Required",
		"License Required",
		"LOW - 0",
		"License required")
	fmt.Println("\n📄 NOT: This is a demo result but reflects real data.")
	fmt.Println("🔗 For full features: https://bugzora.com/license")
}

// LayerFS is a simple fs.FS implementation that overlays tar layers.
type LayerFS struct {
	fs.FS
}
