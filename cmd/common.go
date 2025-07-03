package cmd

import (
	"fmt"
)

func buildTrivyArgs(command, target string) []string {
	args := []string{command, target}

	// Only send --format and --output if Trivy supports the format
	supportedFormats := map[string]bool{
		"table":       true,
		"json":        true,
		"template":    true,
		"sarif":       true,
		"cyclonedx":   true,
		"spdx":        true,
		"spdx-json":   true,
		"github":      true,
		"cosign-vuln": true,
	}

	if outputFormat != "" && supportedFormats[outputFormat] {
		args = append(args, "--format", outputFormat)
	}
	if format != "" && supportedFormats[format] {
		args = append(args, "--format", format)
	}
	if output != "" && supportedFormats[outputFormat] {
		args = append(args, "--output", output)
	}

	if quiet {
		args = append(args, "--quiet")
	}
	if severity != "" {
		args = append(args, "--severity", severity)
	}
	if exitCode != 0 {
		args = append(args, "--exit-code", fmt.Sprintf("%d", exitCode))
	}
	if ignoreUnfixed {
		args = append(args, "--ignore-unfixed")
	}
	if ignorePolicy != "" {
		args = append(args, "--ignore-policy", ignorePolicy)
	}
	if len(skipDirs) > 0 {
		for _, dir := range skipDirs {
			args = append(args, "--skip-dirs", dir)
		}
	}
	if len(skipFiles) > 0 {
		for _, file := range skipFiles {
			args = append(args, "--skip-files", file)
		}
	}
	if listAllPkgs {
		args = append(args, "--list-all-pkgs")
	}
	if offlineScan {
		args = append(args, "--offline-scan")
	}
	if scanners != "" {
		args = append(args, "--scanners", scanners)
	}
	if template != "" {
		args = append(args, "--template", template)
	}
	if configFile != "" {
		args = append(args, "--config", configFile)
	}
	if token != "" {
		args = append(args, "--token", token)
	}
	if proxy != "" {
		args = append(args, "--proxy", proxy)
	}
	if insecure {
		args = append(args, "--insecure")
	}
	if timeout != "" {
		args = append(args, "--timeout", timeout)
	}
	if downloadDBOnly {
		args = append(args, "--download-db-only")
	}
	if debug {
		args = append(args, "--debug")
	}
	if trace {
		args = append(args, "--trace")
	}
	if noProgress {
		args = append(args, "--no-progress")
	}
	if skipUpdate {
		args = append(args, "--skip-update")
	}
	if skipDBUpdate {
		args = append(args, "--skip-db-update")
	}
	if skipPolicyUpdate {
		args = append(args, "--skip-policy-update")
	}
	if securityChecks != "" && securityChecks != "vuln" {
		args = append(args, "--security-checks", securityChecks)
	}
	if compliance != "" {
		args = append(args, "--compliance", compliance)
	}
	if trivyPolicy != "" {
		args = append(args, "--policy", trivyPolicy)
	}
	if len(namespaces) > 0 {
		for _, ns := range namespaces {
			args = append(args, "--namespaces", ns)
		}
	}
	if len(severities) > 0 {
		for _, sev := range severities {
			args = append(args, "--severities", sev)
		}
	}
	if len(ignoreIDs) > 0 {
		for _, id := range ignoreIDs {
			args = append(args, "--ignore-ids", id)
		}
	}
	if ignoreFile != "" {
		args = append(args, "--ignore-file", ignoreFile)
	}
	if includeDevDeps {
		args = append(args, "--include-dev-deps")
	}
	if skipJavaDB {
		args = append(args, "--skip-java-db")
	}
	if light {
		args = append(args, "--light")
	}
	if autoRefresh {
		args = append(args, "--auto-refresh")
	}
	if refresh {
		args = append(args, "--refresh")
	}
	if onlyUpdate != "" {
		args = append(args, "--only-update", onlyUpdate)
	}

	return args
}
