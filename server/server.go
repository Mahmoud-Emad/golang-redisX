package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/Mahmoud-Emad/redisX/resp"
)

func handleConnection(conn net.Conn, server *RespServer) {
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

		CheckCommand(conn, server, command, args)
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
		s.SetClient(conn.RemoteAddr().String())
		terminal.onConnect(s.client.remoteAddress)
		go handleConnection(conn, s)
	}
}

func CheckCommand(conn net.Conn, server *RespServer, command string, args []resp.Value) {
	for i := 0; i < len(server.Commands); i++ {
		if command == server.Commands[i].Name {
			response := server.Commands[i].CallableFunction()
			value, err := resp.ParseRESPValue(response)
			if err != nil {
				// Handle the error.
			}
			conn.Write([]byte(value))
		} else {
			conn.Write([]byte("+Unknown Command: try to register it.\r\n"))
		}
	}
}
