package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var (
	countPtr = flag.Int("c", 0, "Number of ICMP packets to be sent (count)")
	ttlPtr = flag.Int("t", 64, "Initial Time-To-Live (TTL) value of the ICMP packet")
	defaultMessage = "Hii, How are you today?, fine?"
)

func SendingPingMessage(IPTarget string, IPType string, Message string, sequence int, initialTTL int) (time.Duration, error) {
  packetconn, err := icmp.ListenPacket(IPType+":icmp", "0.0.0.0") 
  if err != nil {
    return 0, fmt.Errorf("failed to create icmp socket (try run as root/administrator): %v", err)
  }
  defer packetconn.Close()
  msg := icmp.Message{
    Type: ipv4.ICMPTypeEcho,
    Code: 0,
    Body: &icmp.Echo{
      ID:   os.Getpid() & 0xffff,
      Seq:  sequence,
      Data: []byte(Message),
    },
  }
  b, err := msg.Marshal(nil)
  if err != nil {
    return 0, fmt.Errorf("failed to marshal the message: %v", err)
  }
  targetAddr, err := net.ResolveIPAddr(IPType, IPTarget)
  if err != nil {
    return 0, fmt.Errorf("the target ip address is invalid.: %v", err)
  }
  startTime := time.Now()
  if _, err := packetconn.WriteTo(b, targetAddr); err != nil {
    return 0, fmt.Errorf("failed to send icmp packet: %v", err)
  }
  timeoutDuration := 4 * time.Second
  if err := packetconn.SetReadDeadline(time.Now().Add(timeoutDuration)); err != nil {
    return 0, fmt.Errorf("failed to set a reading time limit: %v", err)
  }
  rb := make([]byte, 1500)
  _, peer, err := packetconn.ReadFrom(rb)

  if err != nil {
    if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
      return 0, fmt.Errorf("timeout (target %s)", IPTarget)
    }
    return 0, fmt.Errorf("error when reading replies: %v", err)
  }
  if peer.String() != IPTarget {
    return 0, fmt.Errorf("unexpected response: %s", peer.String())
  }
  duration := time.Since(startTime)
  return duration, nil
}

func main() {
	flag.Parse() 

	count := *countPtr
	initialTTL := *ttlPtr
	nonFlagArgs := flag.Args()
	
	typeIC := "help"
	targetIC := ""
	messageIC := defaultMessage

	fmt.Print(`
  _         _    ___ _ 
 | |   __ _| |__| _ (_)
 | |__/ _` + "`" + ` | '_ \  _/ |
 |____\__,_|_.__/_| |_|`+"\n\n\x1b[34mLabel ping!\x1b[0m\nAdd a message to label the ICMP packet transmission!\n\n")

	if len(nonFlagArgs) == 0 {
		fmt.Print("Usage   : \x1b[33mlabpi\x1b[0m \x1b[90m[options]\x1b[0m \x1b[35m<Target>\x1b[0m \x1b[37m[Message]\x1b[0m\n")
		fmt.Printf("Example : labpi -c 10 -t 128 192.168.1.1 \"Hii, this me!\"\n")
		fmt.Printf("          labpi 192.168.1.1 \"Hii, this me!\"\n")
		fmt.Printf("          labpi 192.168.1.1\n\n")
		fmt.Print("Default :\n")
		fmt.Printf("• \x1b[90mCount Loop\x1b[0m  : %d \n• \x1b[90mDefault TTL\x1b[0m : %d\n\nRawr uwu nyaw meow chan ehe kawai~\n\n", count, initialTTL)
		os.Exit(0)
	}

	targetArg := nonFlagArgs[0]
	ip := net.ParseIP(targetArg)
	
	if ip == nil {
		fmt.Printf("[\x1b[31mCatch you!\x1b[0m]: \"%s\" not a valid IP address.\n\n", targetArg)
		os.Exit(1)
	}

	targetIC = targetArg
	if ip.To4() != nil {
		typeIC = "ip4"
	} else {
		typeIC = "ip6"
	}
	
	if len(nonFlagArgs) > 1 {
		messageIC = nonFlagArgs[1]
	}

	fmt.Printf("[\x1b[33mTrying!\x1b[0m]: Sending icmp to \"%s\"...\n", targetArg)

	unlimitedLoop := false
	if count == 0 {
		unlimitedLoop = true
	}
	seq := 1
	for i := 0; i < count || unlimitedLoop; i++ {
    duration, err := SendingPingMessage(targetIC, typeIC, messageIC, seq, initialTTL)
    if err != nil {
      fmt.Printf("[\x1b[31mCatch you!\x1b[0m]: Seq=%d %v\n", seq, err)
    } else {
      fmt.Printf("[\x1b[32mWorks!\x1b[0m]: Seq=%d reply from %s, time=%s\n", seq,targetIC, duration.Round(time.Millisecond))
    }
    seq++
    if unlimitedLoop || i < count-1 {
      time.Sleep(1 * time.Second)
    }
    if unlimitedLoop {
      i = 0 
    }
  }
}