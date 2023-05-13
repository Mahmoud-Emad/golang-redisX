package server

import (
	"github.com/google/uuid"
)

// Server helpers functions.

// Init function, will initialize an instance of the server object, then can interacting with all it's items.
// ** Functions **
func Init(host string, port string, network string) RespServer {
	return RespServer{
		id:       uuid.New().String(),
		host:     host,
		Commands: []Command{},
		network:  network,
		port:     port,
	}
}

// Commands helpers functions.

// RegisterCommand, used to register a new command inside the server instance, then you can access it and all it's items.
func (s *RespServer) RegisterCommand(name string, withArgs bool, callableFunction func() string) Command {
	command := Command{
		Name:             name,
		WithArgs:         withArgs,
		CallableFunction: callableFunction,
	}
	s.Commands = append(s.Commands, command)
	return command
}

// This function used for setting the client inside the server, to make all client accessible.
func (s *RespServer) SetClient(address string) *RespServer {
	client := ClientConnection{
		id:            uuid.New().String(),
		remoteAddress: address,
	}
	s.client = &client
	return s
}
