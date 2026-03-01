# piped-go

A lightweight, zero-dependency network traffic monitor written in **Go**. It hooks directly into the Linux kernel to capture raw Ethernet frames without using `libpcap` or any external C libraries.

## 🚀 How it Works
`piped-go` utilizes **Raw Sockets** (`AF_PACKET`). Most programs wait for the OS to process network data; this program "pipes" into the data stream before the OS even sorts it, allowing you to see:
- **Incoming & Outgoing** traffic.
- **Layer 2 (Ethernet)** details like MAC addresses.
- **Full Payloads** of every packet.

## 🛠 Features
- **Zero Libs**: Built using only the [Go Standard Library](https://pkg.go.dev).
- **Direct Access**: Uses `syscall` to talk directly to the Linux kernel.
- **Full Transparency**: No headers are stripped; you see the raw data as it hits the wire.

## 📦 Installation & Build
Since it's pure Go, you just need the compiler:

```bash
# Build the binary
go build -o piped-go main.go
