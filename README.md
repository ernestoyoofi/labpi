# LabPi - Label Ping

`labpi` is a specialized ping utility designed to allow users to embed a custom message within the ICMP packet transmission. This feature enables easy labeling or identification of the packets sent, which can be useful for network debugging, tracking specific test packets, or just adding a fun signature!

## 📦 Installation

Check the [release in this repository](https://github.com/ernestoyoofi/labpi/releases), then select according to your architecture and operating system. If you can't find it, you can build it yourself by cloning this repository.

## 🚀 Usage

The basic syntax for using labpi is straightforward:

```bash
labpi [options] <Target> [Message]
```

**Parameters :**

- `[options]` Optional flags to control the ping behavior.
- `<Target>` Required. The IP address or hostname of the target machine.
- `[Message]` Optional. The custom message to embed in the ICMP packet.

**Options :**

- `-c <Count>` Stop after sending (and receiving) Count ECHO_RESPONSE packets.
- `-t <TTL>` Set the IP Time To Live.
