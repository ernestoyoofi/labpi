package cli

import (
	"fmt"
	"label-ping/packet"
	"time"
)

// PrintHeader prints the LabPi header
func PrintHeader() {
	fmt.Print(`
  _         _    ___ _ 
 | |   __ _| |__| _ (_)
 | |__/ _` + "`" + ` | '_ \  _/ |
 |____\__,_|_.__/_| |_|` + "\n\n")
	fmt.Print("\x1b[34mLabel ping!\x1b[0m\nAdd a message to label the ICMP packet transmission!\n\n")
}

// PrintUsage prints usage information
func PrintUsage() {
	fmt.Print("Usage   : \x1b[33mlabpi\x1b[0m \x1b[90m[options]\x1b[0m \x1b[35m<Target>\x1b[0m \x1b[37m[Message]\x1b[0m\n")
	fmt.Printf("Example : labpi -c 10 -t 128 192.168.1.1 \"Hii, this me!\"\n")
	fmt.Printf("          labpi 192.168.1.1 \"Hii, this me!\"\n")
	fmt.Printf("          labpi 192.168.1.1\n\n")
	fmt.Print("Options:\n")
	fmt.Print("  -c <count>  Stop after sending COUNT ECHO_RESPONSE packets (0 = unlimited)\n")
	fmt.Print("  -t <ttl>    Set the IP Time To Live (default 64)\n\n")
}

// Run executes the ping sequence with the given config
func Run(cfg *Config) int {
	PrintHeader()

	fmt.Printf("[\x1b[33mTrying!\x1b[0m]: Sending icmp to \"%s\"...\n", cfg.Target)

	seq := 1
	successCount := 0
	failureCount := 0

	// Determine if unlimited loop
	unlimited := cfg.Count == 0

	for i := 0; i < cfg.Count || unlimited; i++ {
		result := packet.SendICMP(cfg.Target, cfg.Message, seq, cfg.TTL, 4*time.Second)

		if result.Error != nil {
			fmt.Printf("[\x1b[31mCatch you!\x1b[0m]: Seq=%d %v\n", seq, result.Error)
			failureCount++
		} else {
			fmt.Printf("[\x1b[32mWorks!\x1b[0m]: Seq=%d reply from %s, time=%v\n", seq, cfg.Target, result.Duration.Round(time.Millisecond))
			successCount++
		}

		seq++

		// Sleep between packets, but not after the last one
		if unlimited || i < cfg.Count-1 {
			time.Sleep(1 * time.Second)
		}

		// Reset counter for unlimited loop
		if unlimited {
			i = -1
		}
	}

	fmt.Printf("\n[\x1b[36mStats!\x1b[0m]: %d sent, %d received, %d lost\n\n", seq-1, successCount, failureCount)
	return 0
}
