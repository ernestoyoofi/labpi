package packet

import "testing"

func TestValidateTTL(t *testing.T) {
	tests := []struct {
		name    string
		ttl     int
		wantErr bool
	}{
		{"valid min TTL", 0, false},
		{"valid max TTL", 255, false},
		{"valid mid TTL", 128, false},
		{"negative TTL", -1, true},
		{"TTL 256", 256, true},
		{"TTL 1000", 1000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTTL(tt.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTTL(%d) error = %v, wantErr %v", tt.ttl, err, tt.wantErr)
			}
		})
	}
}

func TestValidateMessage(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		ipType  string
		wantErr bool
	}{
		{"empty message", "", "ip4", false},
		{"short message", "hello", "ip4", false},
		{"max valid message", string(make([]byte, 65507)), "ip4", false},
		{"oversized message", string(make([]byte, 65508)), "ip4", true},
		{"ipv6 short message", "hello", "ip6", false},
		{"ipv6 max message", string(make([]byte, 65507)), "ip6", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMessage(tt.msg, tt.ipType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
