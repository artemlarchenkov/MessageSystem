package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
	Address        string
	MessageHandler func(string)
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	fmt.Println("Receiver started")

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept error", err)
			}

			go s.handleConnection(conn)
		}
	}()

	return nil
}
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		msg = msg[:len(msg)-1]
		s.MessageHandler(msg)
		fmt.Println("Received TCP:", msg)
	}
}
