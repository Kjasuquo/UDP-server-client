package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	add := net.UDPAddr{
		Port: 1357,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.DialUDP("udp", &add, &addr)
	if err != nil {
		fmt.Println("error occurred")
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error occurred ")
		}

		conn.Write([]byte(message))
	}
}
