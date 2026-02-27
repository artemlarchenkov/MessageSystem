package main

import "messagesystem/internal/api"

func main() {
	a := &api.API{
		TCPAddr: "receiver:9000",
	}

	r := a.Routes()
	r.Run(":8081") // REST API sender
}
