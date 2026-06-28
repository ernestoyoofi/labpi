package packet

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// Result holds the result of an ICMP echo request
type Result struct {
	Duration time.Duration
	Seq      int
	Error    error
}

// SendICMP sends an ICMP echo request and waits for a reply, with TCP fallback
func SendICMP(target string, message string, seq int, ttl int, timeout time.Duration) Result {
	ip := net.ParseIP(target)
	if ip == nil {
		return Result{Seq: seq, Error: fmt.Errorf("invalid IP address: %s", target)}
	}

	// Try native ICMP first
	result := trySendICMP(target, message, seq, ttl, timeout)
	if result.Error == nil {
		return result
	}

	// Fallback to TCP ping when ICMP unavailable (no privileges)
	return sendTCPPing(target, seq, timeout)
}

func trySendICMP(target string, message string, seq int, ttl int, timeout time.Duration) Result {
	ip := net.ParseIP(target)
	if ip == nil {
		return Result{Seq: seq, Error: fmt.Errorf("invalid IP address: %s", target)}
	}

	var network string
	if ip.To4() != nil {
		network = "ip4:icmp"
	} else {
		network = "ip6:icmp"
	}

	conn, err := icmp.ListenPacket(network, "0.0.0.0")
	if err != nil {
		// Try alternative bind address for IPv6
		if network == "ip6:icmp" {
			conn, err = icmp.ListenPacket(network, "::")
			if err != nil {
				return Result{Seq: seq, Error: err}
			}
		} else {
			return Result{Seq: seq, Error: err}
		}
	}
	defer conn.Close()

	// Set TTL if specified
	if ttl > 0 {
		if ip.To4() != nil {
			if c := conn.IPv4PacketConn(); c != nil {
				c.SetTTL(ttl)
			}
		} else {
			if c := conn.IPv6PacketConn(); c != nil {
				c.SetHopLimit(ttl)
			}
		}
	}

	// Build ICMP message
	var msg icmp.Message
	if ip.To4() != nil {
		msg = icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  seq,
				Data: []byte(message),
			},
		}
	} else {
		msg = icmp.Message{
			Type: ipv6.ICMPTypeEchoRequest,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  seq,
				Data: []byte(message),
			},
		}
	}

	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		return Result{Seq: seq, Error: err}
	}

	targetAddr, err := net.ResolveIPAddr("ip", target)
	if err != nil {
		return Result{Seq: seq, Error: err}
	}

	start := time.Now()
	_, err = conn.WriteTo(msgBytes, targetAddr)
	if err != nil {
		return Result{Seq: seq, Error: err}
	}

	// Set read deadline
	conn.SetReadDeadline(time.Now().Add(timeout))

	reply := make([]byte, 1500)
	_, peer, err := conn.ReadFrom(reply)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return Result{Seq: seq, Error: fmt.Errorf("timeout")}
		}
		return Result{Seq: seq, Error: err}
	}

	duration := time.Since(start)

	if peer.String() != target {
		return Result{Seq: seq, Error: fmt.Errorf("unexpected reply from %s", peer.String())}
	}

	return Result{Duration: duration, Seq: seq, Error: nil}
}

// sendTCPPing fallback untuk ketika ICMP tidak available (no privileges)
func sendTCPPing(target string, seq int, timeout time.Duration) Result {
	ports := []string{"80", "443"}
	targetAddr := target
	if net.ParseIP(target).To4() == nil && net.ParseIP(target) != nil {
		targetAddr = "[" + target + "]"
	}

	for _, port := range ports {
		start := time.Now()
		conn, err := net.DialTimeout("tcp", targetAddr+":"+port, timeout)
		if err == nil {
			conn.Close()
			return Result{Duration: time.Since(start), Seq: seq, Error: nil}
		}
	}

	hint := ""
	switch runtime.GOOS {
	case "linux":
		hint = " (try: sudo sysctl -w net.ipv4.ping_group_range='0 2147483647')"
	case "darwin":
		hint = " (try: sudo labpi)"
	case "windows":
		hint = " (run as Administrator)"
	}

	return Result{Seq: seq, Error: fmt.Errorf("ICMP unavailable - TCP fallback failed%s", hint)}
}
