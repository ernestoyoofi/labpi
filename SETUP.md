# LabPi - Setup & Unprivileged ICMP

## Installation

### Binary

```bash
# Download from GitHub Releases
wget https://github.com/ernestoyoofi/labpi/releases/download/v1.0.0/labpi-linux-amd64
chmod +x labpi-linux-amd64
./labpi-linux-amd64 8.8.8.8
```

### Package Managers

**Debian/Ubuntu:**
```bash
wget https://github.com/ernestoyoofi/labpi/releases/download/v1.0.0/labpi-amd64.deb
sudo apt install ./labpi-amd64.deb
labpi 8.8.8.8
```

**Fedora/RHEL/CentOS:**
```bash
wget https://github.com/ernestoyoofi/labpi/releases/download/v1.0.0/labpi-1.0.0-1.x86_64.rpm
sudo dnf install ./labpi-*.rpm
labpi 8.8.8.8
```

**Arch Linux:**
```bash
git clone https://github.com/ernestoyoofi/labpi.git
cd labpi
makepkg -si
labpi 8.8.8.8
```

### Docker

```bash
docker run --rm ernestoyoofi/labpi 8.8.8.8
```

## Quick Start

```bash
# Without privileges (uses TCP fallback on Linux/macOS if needed)
labpi 8.8.8.8
labpi -c 5 192.168.1.1 "custom message"
```

## Platform-Specific Setup

### Linux - Enable Unprivileged ICMP (Recommended)

To use ICMP directly without `sudo`:

```bash
# Temporary (until reboot)
sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"

# Permanent
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

# Option 2: Grant setcap (if supported)
sudo setcap cap_net_raw=ep ./labpi
```

### Windows

ICMP usually works without admin on modern Windows. If needed:
- Run Command Prompt as Administrator
- Run labpi directly

## Fallback Mechanism

If ICMP is unavailable (no privileges), LabPi automatically falls back to **TCP ping** on ports 80/443.

- **ICMP available**: Direct ICMP echo
- **ICMP denied**: TCP connect to port 80 or 443
- **Both fail**: Error with helpful setup command

## Usage

```bash
labpi [options] <Target> [Message]

Options:
  -c <count>  Number of pings (0 = unlimited, default)
  -t <ttl>    Time To Live (default 64)

Examples:
  labpi 8.8.8.8
  labpi -c 5 google.com "Hello"
  labpi -t 128 2001:4860:4860::8888
```

## Cross-Platform Support

Works on: Windows, Linux, macOS, BSD, Android, FreeBSD, OpenBSD, NetBSD, Plan9, Solaris
