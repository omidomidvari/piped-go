package main

import (
	"fmt"
	"log"
	"net"
	"syscall"
)

func main() {
	// Create a raw socket to capture all ethernet frames (ETH_P_ALL = 0x0300)
	// Requires root/CAP_NET_RAW privileges.
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_ALL)))
	if err != nil {
		log.Fatalf("Error creating socket: %v", err)
	}
	defer syscall.Close(fd)

	fmt.Println("piped-go: Monitoring all network data (press Ctrl+C to stop)...")

	buffer := make([]byte, 65535)
	for {
		// Read raw data directly from the socket
		n, addr, err := syscall.Recvfrom(fd, buffer, 0)
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}

		// Basic output: Timestamp, byte count, and source address info
		fmt.Printf("[%s] Received %d bytes from %+v\n", 
			net.HardwareAddr(buffer[6:12]), // Source MAC address from Ethernet header
			n, 
			addr,
		)
	}
}

// htons converts a short integer from host byte order to network byte order.
func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}
