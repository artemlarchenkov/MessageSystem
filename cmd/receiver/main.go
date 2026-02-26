package main

import (
	"messagesystem/internal/api"
	"messagesystem/internal/storage"
	"messagesystem/internal/tcp"
)

func main() {
	st, err := storage.New("messages.db")
	if err != nil {
		panic(err)
	}

	server := &tcp.Server{
		Address: ":9000",
		MessageHandler: func(msg string) {
			st.SaveMessage(msg)
		},
	}
	server.Start()

	a := &api.API{
		Storage: st,
		TCPAddr: "localhost:9000",
	}
	r := a.Routes()
	r.Run(":8080")
}
