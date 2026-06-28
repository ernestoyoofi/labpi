package cli

import (
	"testing"
)

func TestParseArgs_MissingTarget(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{})
	if err == nil {
		t.Error("expected error for missing target, got nil")
	}
	if cfg != nil {
		t.Error("expected nil config, got non-nil")
	}
}

func TestParseArgs_InvalidIP(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{"invalid-ip"})
	if err == nil {
		t.Error("expected error for invalid IP, got nil")
	}
	if cfg != nil {
		t.Error("expected nil config, got non-nil")
	}
}

func TestParseArgs_ValidWithDefaults(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{"192.168.1.1"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if cfg == nil {
		t.Error("expected non-nil config, got nil")
	}
	if cfg.Target != "192.168.1.1" {
		t.Errorf("expected target 192.168.1.1, got %s", cfg.Target)
	}
	if cfg.Count != 0 {
		t.Errorf("expected count 0, got %d", cfg.Count)
	}
	if cfg.TTL != 64 {
		t.Errorf("expected TTL 64, got %d", cfg.TTL)
	}
	if cfg.Message == "" {
		t.Error("expected default message, got empty")
	}
}

func TestParseArgs_WithFlags(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{"-c", "5", "-t", "128", "8.8.8.8", "hello"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if cfg == nil {
		t.Error("expected non-nil config, got nil")
	}
	if cfg.Count != 5 {
		t.Errorf("expected count 5, got %d", cfg.Count)
	}
	if cfg.TTL != 128 {
		t.Errorf("expected TTL 128, got %d", cfg.TTL)
	}
	if cfg.Message != "hello" {
		t.Errorf("expected message 'hello', got %s", cfg.Message)
	}
	if cfg.Target != "8.8.8.8" {
		t.Errorf("expected target 8.8.8.8, got %s", cfg.Target)
	}
}

func TestParseArgs_InvalidTTL(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{"-t", "256", "192.168.1.1"})
	if err == nil {
		t.Error("expected error for invalid TTL, got nil")
	}
	if cfg != nil {
		t.Error("expected nil config, got non-nil")
	}
}

func TestParseArgs_IPv6Target(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{"2001:4860:4860::8888"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if cfg == nil {
		t.Error("expected non-nil config, got nil")
	}
	if cfg.Target != "2001:4860:4860::8888" {
		t.Errorf("expected IPv6 target, got %s", cfg.Target)
	}
}

func TestParseArgs_MissingFlagValue(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{"-c", "192.168.1.1"})
	if err == nil {
		t.Error("expected error for missing -c value, got nil")
	}
	if cfg != nil {
		t.Error("expected nil config, got non-nil")
	}
}

func TestParseArgs_OnlyTarget(t *testing.T) {
	cfg, err := ParseArgsWithArgs([]string{"127.0.0.1"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if cfg.Target != "127.0.0.1" {
		t.Errorf("expected target 127.0.0.1, got %s", cfg.Target)
	}
}

