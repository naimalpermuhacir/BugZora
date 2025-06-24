/*
Copyright © 2025 BugZora <bugzora@bugzora.dev>
*/
package main

import "bugzora/cmd"

// Version information - set by GoReleaser
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	cmd.Execute()
}
