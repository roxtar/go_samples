package main

import "net"
import "log"

func main() {
	addr := net.UDPAddr{
		Port: 61555,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	defer conn.Close()
	if err != nil {
		log.Fatal("Error: %v", err)
	}
	for {
		buffer := make([]byte, 1024)
		read, _, err := conn.ReadFrom(buffer)

		if err != nil {
			log.Fatal("Error: %v", err)
		}
		println(string(buffer[:read]))
	}
}
