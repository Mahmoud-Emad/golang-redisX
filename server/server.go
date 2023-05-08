package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/Mahmoud-Emad/redisX/resp"
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

		value, err := resp.DecodeRESP(bufio.NewReader(conn))
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("Error decoding RESP: ", err.Error())
			return
		}

		command := value.Array()[0].String()
		args := value.Array()[1:]

		switch command {
		case "ping":
			conn.Write([]byte("+PONG\r\n"))
		case "echo":
			conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0].String()), args[0].String())))
		default:
			conn.Write([]byte("-ERR unknown command '" + command + "'\r\n"))
		}
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
