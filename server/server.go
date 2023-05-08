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
	ln, err := net.Listen(s.network, s.host+":"+s.port)
	if err != nil {
		fmt.Println("Failed to bind to port" + s.port)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	terminal := initTerminal(s)
	terminal.Welcome()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}
