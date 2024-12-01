package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("error closing connection:", err)
		}
	}(conn)

	data := "Hello, server!"
	if _, err := conn.Write([]byte(data)); err != nil {
		fmt.Println("Error sending datagram:", err)
		return
	}

	buf := make([]byte, len(data))
	if _, err := conn.Read(buf); err != nil {
		fmt.Println("Error reading datagram:", err)
		return
	}

	fmt.Println("Received from server:", string(buf))
}
