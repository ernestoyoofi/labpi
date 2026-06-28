package packet

import "fmt"

// ValidateTTL validates TTL is within valid range (0-255)
func ValidateTTL(ttl int) error {
	if ttl < 0 || ttl > 255 {
		return fmt.Errorf("TTL must be between 0 and 255, got %d", ttl)
	}
	return nil
}

// ValidateMessage validates message size based on IP type
// IPv4: max payload ~65507 bytes (65535 - 20 header - 8 ICMP header)
// IPv6: max payload ~65507 bytes (65535 - 40 header - 8 ICMP header)
func ValidateMessage(msg string, ipType string) error {
	msgBytes := len([]byte(msg))
	const maxPayload = 65507

	if msgBytes > maxPayload {
		return fmt.Errorf("message too large: %d bytes (max %d bytes)", msgBytes, maxPayload)
	}
	return nil
}
