/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"bugzora/cmd"

	"github.com/fatih/color"
)

// Version information - set by GoReleaser
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func init() {
	color.NoColor = false // Always enable colored output
}

func main() {
	// Demo mode check
	if os.Getenv("DEMO_MODE") == "true" || isDemoMode() {
		showDemoWarning()
	}

	cmd.Execute()
}

func isDemoMode() bool {
	// Demo mode check - basit kontrol
	return true // Always show demo warning in demo mode
}

func showDemoWarning() {
	fmt.Println("🚨 DEMO MODE")
	fmt.Println("📧 Contact: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}
