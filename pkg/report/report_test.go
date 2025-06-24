package report

import (
	"os"
	"testing"
)

func TestWriteResultsInvalidFormat(t *testing.T) {
	// Test invalid format
	err := WriteResults(nil, "invalid", "test-target")
	if err == nil {
		t.Error("WriteResults should fail with invalid format")
	}
}

func TestWriteResultsTableFormat(t *testing.T) {
	// Test table format (should not create file)
	err := WriteResults(nil, "table", "test-target")
	if err != nil {
		t.Errorf("WriteResults table format failed: %v", err)
	}
}

func TestWriteResultsEmptyResults(t *testing.T) {
	// Test with empty results
	err := WriteResults(nil, "json", "test-target")
	if err != nil {
		t.Errorf("WriteResults with empty results failed: %v", err)
	}

	// Clean up if file was created
	os.Remove("report-test-target.json")
	os.Remove("report-test-target.pdf")
}
