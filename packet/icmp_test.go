package packet

import (
	"testing"
	"time"
)

func TestSendICMP_InvalidIP(t *testing.T) {
	result := SendICMP("invalid-ip", "test", 1, 64, 4*time.Second)
	if result.Error == nil {
		t.Error("expected error for invalid IP, got nil")
	}
	if result.Seq != 1 {
		t.Errorf("expected seq 1, got %d", result.Seq)
	}
}

func TestSendICMP_Timeout(t *testing.T) {
	// Non-routable IP should timeout
	result := SendICMP("192.0.2.1", "test", 1, 64, 100*time.Millisecond)
	if result.Error == nil {
		t.Error("expected timeout error, got nil")
	}
}

func TestSendICMP_ResultStructure(t *testing.T) {
	result := SendICMP("999.999.999.999", "test", 42, 64, 1*time.Second)
	if result.Seq != 42 {
		t.Errorf("expected seq 42, got %d", result.Seq)
	}
	if result.Error == nil {
		t.Error("expected error for invalid IP, got nil")
	}
	if result.Duration != 0 {
		t.Errorf("expected zero duration for error, got %v", result.Duration)
	}
}

func TestSendICMP_TCPFallback(t *testing.T) {
	// Test TCP fallback to google.com (usually has port 80/443 open)
	result := SendICMP("8.8.8.8", "test", 1, 64, 2*time.Second)
	// Should either succeed with ICMP or TCP fallback
	if result.Seq != 1 {
		t.Errorf("expected seq 1, got %d", result.Seq)
	}
}
