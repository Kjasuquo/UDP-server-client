package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error occured")
	}

	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(message)

	}
}
