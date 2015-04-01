package main

import "net"
import "log"

func main() {
	conn, err := net.Dial("udp", "localhost:61555")
	if err != nil {
		log.Fatal("Error: %v", err)
	}
	defer conn.Close()
	str := []byte("test test\n")
	_, err = conn.Write(str)
	if err != nil {
		log.Fatal("Error: %v", err)
	}
}
