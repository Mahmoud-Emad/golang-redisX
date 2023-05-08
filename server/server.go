package server

import (
	"fmt"
	"io"
	"net"
	"os"
)

type RespServer struct {
	host    string
	port    string
	network string
}

func New(host string, port string, network string) RespServer {
	return RespServer{
		host,
		port,
		network,
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		if _, err := conn.Read(buf); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("error reading from client: ", err.Error())
				os.Exit(1)
			}
		}

		// Let's ignore the client's input for now and hardcode a response.
		// We'll implement a proper Redis Protocol parser in later stages.
		conn.Write([]byte("+PONG\r\n"))
	}
}

func (s *RespServer) RunAndWait() {
	terminal := initTerminal(s)
	ln, err := net.Listen(s.network, s.host+":"+s.port)
	if err != nil {
		terminal.RaisError("Failed to bind to port "+s.port, err)
	}
	terminal.Welcome()
	for {
		conn, err := ln.Accept()
		if err != nil {
			terminal.RaisError("Error accepting connection", err)
		}

		go handleConnection(conn)
	}
}
