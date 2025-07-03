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
	color.NoColor = false // Her durumda renkli çıktı
}

func main() {
	// Demo modu kontrolü
	if os.Getenv("DEMO_MODE") == "true" || isDemoMode() {
		showDemoWarning()
	}

	cmd.Execute()
}

func isDemoMode() bool {
	// Demo modu kontrolü - basit kontrol
	return true // Demo modunda her zaman demo uyarısı
}

func showDemoWarning() {
	fmt.Println("🚨 DEMO MODU")
	fmt.Println("📧 İletişim: license@bugzora.com")
	fmt.Println(strings.Repeat("─", 50))
}
