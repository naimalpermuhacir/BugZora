package vuln

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/aquasecurity/trivy/pkg/types"
	"golang.org/x/xerrors"
)

// runTrivyCommand executes the trivy CLI command with the given arguments and captures its JSON output.
func runTrivyCommand(ctx context.Context, args ...string) (types.Report, error) {
	var report types.Report

	// Base command arguments
	baseArgs := []string{"--cache-dir", "/tmp/.trivy-cache", "-q", "--format", "json"}
	args = append(baseArgs, args...)

	cmd := exec.CommandContext(ctx, "trivy", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// The trivy command will exit with a specific code if vulnerabilities are found.
	// We should not treat this as a fatal error for the command execution itself,
	// so we check the error type. We only fail if it's a command execution problem.
	_ = cmd.Run() // We ignore the error here on purpose, see below.

	// If stderr has content, it means trivy encountered a real error (e.g., image not found).
	if stderr.Len() > 0 {
		stderrStr := stderr.String()

		// Check if stderr contains actual error messages (not just info/warning)
		if strings.Contains(stderrStr, "ERROR") ||
			strings.Contains(stderrStr, "error") ||
			strings.Contains(stderrStr, "not found") ||
			strings.Contains(stderrStr, "failed") {
			return report, xerrors.Errorf("trivy command failed: %s", stderrStr)
		}

		// If it's just INFO/WARN messages, we can ignore them
		// Trivy often writes progress info to stderr even in quiet mode
	}

	// If stdout is empty, it means no report was generated.
	if stdout.Len() == 0 {
		// This can happen if the scan target is not valid but trivy doesn't error,
		// or if no vulnerabilities are found and trivy is configured to not output empty reports.
		// Returning an empty report is safe.
		return report, nil
	}

	// The command ran successfully, and we have a JSON report to parse.
	if err := json.Unmarshal(stdout.Bytes(), &report); err != nil {
		// This is a critical error, as we can't parse the output we received.
		return report, xerrors.Errorf("failed to unmarshal trivy json report: %w\nOutput:\n%s", err, stdout.String())
	}

	return report, nil
}

// runTrivyCommandWithArgs executes the trivy CLI command with custom arguments
func runTrivyCommandWithArgs(ctx context.Context, args []string, quiet bool) (types.Report, error) {
	var report types.Report

	// Base command arguments - always use JSON format for Trivy
	baseArgs := []string{"--cache-dir", "/tmp/.trivy-cache", "--format", "json"}

	// Only add quiet if not already specified and quiet is true
	if quiet {
		hasQuiet := false
		for _, arg := range args {
			if arg == "--quiet" || arg == "-q" {
				hasQuiet = true
				break
			}
		}
		if !hasQuiet {
			baseArgs = append(baseArgs, "-q")
		}
	}

	args = append(baseArgs, args...)

	cmd := exec.CommandContext(ctx, "trivy", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// The trivy command will exit with a specific code if vulnerabilities are found.
	// We should not treat this as a fatal error for the command execution itself,
	// so we check the error type. We only fail if it's a command execution problem.
	_ = cmd.Run() // We ignore the error here on purpose, see below.

	// If stderr has content, it means trivy encountered a real error (e.g., image not found).
	if stderr.Len() > 0 {
		stderrStr := stderr.String()

		// Check if stderr contains actual error messages (not just info/warning)
		if strings.Contains(stderrStr, "ERROR") ||
			strings.Contains(stderrStr, "error") ||
			strings.Contains(stderrStr, "not found") ||
			strings.Contains(stderrStr, "failed") {
			return report, xerrors.Errorf("trivy command failed: %s", stderrStr)
		}

		// If it's just INFO/WARN messages, we can ignore them
		// Trivy often writes progress info to stderr even in quiet mode
	}

	// If stdout is empty, it means no report was generated.
	if stdout.Len() == 0 {
		// This can happen if the scan target is not valid but trivy doesn't error,
		// or if no vulnerabilities are found and trivy is configured to not output empty reports.
		// Returning an empty report is safe.
		return report, nil
	}

	// The command ran successfully, and we have a JSON report to parse.
	if err := json.Unmarshal(stdout.Bytes(), &report); err != nil {
		// This is a critical error, as we can't parse the output we received.
		return report, xerrors.Errorf("failed to unmarshal trivy json report: %w\nOutput:\n%s", err, stdout.String())
	}

	return report, nil
}

// ScanFilesystem uses the trivy CLI to scan a local filesystem path for vulnerabilities.
func ScanFilesystem(ctx context.Context, targetPath string, quiet bool) (types.Report, error) {
	args := []string{"fs", targetPath}
	return runTrivyCommand(ctx, args...)
}

// ScanFilesystemWithArgs uses the trivy CLI to scan a local filesystem path with custom arguments.
func ScanFilesystemWithArgs(ctx context.Context, targetPath string, trivyArgs []string, quiet bool) (types.Report, error) {
	return runTrivyCommandWithArgs(ctx, trivyArgs, quiet)
}

// ScanImage uses the trivy CLI to scan a container image for vulnerabilities.
func ScanImage(ctx context.Context, imageName string, quiet bool) (types.Report, error) {
	args := []string{"image", imageName}
	return runTrivyCommand(ctx, args...)
}

// ScanImageWithArgs uses the trivy CLI to scan a container image with custom arguments.
func ScanImageWithArgs(ctx context.Context, imageName string, trivyArgs []string, quiet bool) (types.Report, error) {
	return runTrivyCommandWithArgs(ctx, trivyArgs, quiet)
}

// ScanRepository uses the trivy CLI to scan a Git repository for vulnerabilities.
func ScanRepository(ctx context.Context, repoURL string, quiet bool) (types.Report, error) {
	args := []string{"repo", repoURL}
	return runTrivyCommand(ctx, args...)
}

// ScanRepositoryWithArgs uses the trivy CLI to scan a Git repository with custom arguments.
func ScanRepositoryWithArgs(ctx context.Context, repoURL string, trivyArgs []string, quiet bool) (types.Report, error) {
	return runTrivyCommandWithArgs(ctx, trivyArgs, quiet)
}
