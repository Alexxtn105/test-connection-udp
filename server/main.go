package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening on :8080...")

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			fmt.Println("Error reading datagram:", err)
			continue
		}

		if _, err := conn.WriteTo(buf[:n], addr); err != nil {
			fmt.Println("Error writing datagram:", err)
		}
	}
}
