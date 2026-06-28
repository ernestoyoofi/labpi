package cli

import (
	"testing"
)

func TestPrintHeader(t *testing.T) {
	// Just ensure it doesn't panic
	PrintHeader()
}

func TestPrintUsage(t *testing.T) {
	// Just ensure it doesn't panic
	PrintUsage()
}

func TestRun_WithLimitedCount(t *testing.T) {
	cfg := &Config{
		Target:  "127.0.0.1",
		Message: "test",
		Count:   1,
		TTL:     64,
	}
	// Should return 0 (success exit code)
	code := Run(cfg)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
}
