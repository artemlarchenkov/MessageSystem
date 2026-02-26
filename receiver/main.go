package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	fmt.Println("Receiver started...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		fmt.Println("Received:", string(buf[:n]))
	}
}
