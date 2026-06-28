package cli

import (
	"fmt"
	"label-ping/packet"
	"net"
	"os"
	"strconv"
)

// Config holds parsed CLI configuration
type Config struct {
	Target  string
	Message string
	Count   int
	TTL     int
}

// ParseArgs parses command line arguments and returns Config
func ParseArgs() (*Config, error) {
	return ParseArgsWithArgs(os.Args[1:])
}

// ParseArgsWithArgs parses provided arguments (useful for testing)
func ParseArgsWithArgs(args []string) (*Config, error) {
	count := 0
	ttl := 64
	var target, message string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-c":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("flag -c requires an argument")
			}
			i++
			val, err := strconv.Atoi(args[i])
			if err != nil {
				return nil, fmt.Errorf("invalid count value: %v", err)
			}
			count = val
		case "-t":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("flag -t requires an argument")
			}
			i++
			val, err := strconv.Atoi(args[i])
			if err != nil {
				return nil, fmt.Errorf("invalid TTL value: %v", err)
			}
			ttl = val
		default:
			if target == "" {
				target = args[i]
			} else if message == "" {
				message = args[i]
			}
		}
	}

	if target == "" {
		return nil, fmt.Errorf("target IP address is required")
	}

	// Validate target is valid IP
	if net.ParseIP(target) == nil {
		return nil, fmt.Errorf("invalid IP address: %s", target)
	}

	// Validate TTL
	if err := packet.ValidateTTL(ttl); err != nil {
		return nil, err
	}

	// Default message
	if message == "" {
		message = "Hii, How are you today?, fine?"
	}

	// Validate message
	if err := packet.ValidateMessage(message, ""); err != nil {
		return nil, err
	}

	return &Config{
		Target:  target,
		Message: message,
		Count:   count,
		TTL:     ttl,
	}, nil
}
