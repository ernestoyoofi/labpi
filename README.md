# LabPi - Label Ping

`labpi` is a specialized ping utility designed to allow users to embed a custom message within the ICMP packet transmission. This feature enables easy labeling or identification of the packets sent, which can be useful for network debugging, tracking specific test packets, or just adding a fun signature!

## Features

- **Custom ICMP messages** - Embed labels in packets for identification
- **Unprivileged ICMP** - Works without `sudo` on Linux with proper setup (or automatic TCP fallback)
- **Cross-platform** - Windows, Linux, macOS, BSD, Android, FreeBSD, OpenBSD, NetBSD, Plan9, Solaris
- **IPv4 & IPv6 support** - Full dual-stack networking
- **TCP fallback** - Automatically falls back to TCP ping when ICMP unavailable
- **Clean architecture** - Modular design with comprehensive tests

## Installation

### Pre-built Binaries

Download from [GitHub Releases](https://github.com/ernestoyoofi/labpi/releases) for your platform.

### Linux Package Managers

**Debian/Ubuntu:**
```bash
sudo apt install ./labpi-amd64.deb
```

**Fedora/RHEL/CentOS:**
```bash
sudo dnf install ./labpi-*.rpm
```

**Arch Linux:**
```bash
git clone https://github.com/ernestoyoofi/labpi.git
cd labpi
makepkg -si
```

### Docker

```bash
docker run --rm ernestoyoofi/labpi 8.8.8.8
docker run --rm ernestoyoofi/labpi -c 5 192.168.1.1 "hello"
```

### Build from Source

```bash
git clone https://github.com/ernestoyoofi/labpi.git
cd labpi
go build -o labpi
./labpi 8.8.8.8
```

## Usage

```bash
labpi [options] <Target> [Message]
```

**Parameters:**
- `<Target>` - IP address or hostname (IPv4 or IPv6)
- `[Message]` - Custom message to embed in packet (optional, default: "Hii, How are you today?, fine?")

**Options:**
- `-c <count>` - Stop after COUNT pings (0 = unlimited, default 0)
- `-t <ttl>` - Set Time To Live (0-255, default 64)

**Examples:**
```bash
# Simple ping
labpi 8.8.8.8

# With custom message
labpi 192.168.1.1 "Hello from labpi"

# 5 pings with TTL 128
labpi -c 5 -t 128 google.com "custom label"

# IPv6 support
labpi 2001:4860:4860::8888 "IPv6 test"
```

## Setup for Unprivileged ICMP

### Linux - Enable Unprivileged Ping

To use ICMP without `sudo`, run once:
```bash
sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"
```

Or make it permanent:
```bash
echo "net.ipv4.ping_group_range = 0 2147483647" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```

Alternative: Grant capability to binary
```bash
sudo setcap cap_net_raw=ep ./labpi
```

### macOS

```bash
# Option 1: Use sudo
sudo labpi 192.168.1.1

# Option 2: Grant capability (if supported)
sudo setcap cap_net_raw=ep ./labpi
```

### Windows

ICMP typically works without admin. If needed, run Command Prompt as Administrator.

## Automatic TCP Fallback

If ICMP socket is unavailable (no privileges), LabPi automatically attempts TCP ping to port 80/443:
- ICMP available → Direct ICMP echo response
- ICMP denied (no privileges) → TCP connect ping (port 80/443)
- Both fail → Clear error with setup instructions

This ensures basic connectivity testing works on all environments without privilege escalation.

## Architecture

LabPi is built with clean separation of concerns:
- **`packet/`** - ICMP/TCP sender abstraction with cross-platform support
- **`cli/`** - Argument parsing, output formatting, main loop orchestration
- **`main.go`** - Minimal entry point (19 lines)

See [SETUP.md](SETUP.md) for detailed platform-specific information.

## 🐳 Docker

Build and run with Docker:

```bash
docker build -t labpi .
docker run --rm labpi 8.8.8.8
docker run --rm labpi -c 5 -t 128 google.com "Docker test"
```

Or use pre-built image:
```bash
docker run --rm ernestoyoofi/labpi 8.8.8.8
```
