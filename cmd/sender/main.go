package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		msg := "Hello from sender\n"
		conn.Write([]byte(msg))
		fmt.Println("Sent:", msg)
	}
}
