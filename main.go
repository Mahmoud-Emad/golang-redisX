package main

import (
	"github.com/Mahmoud-Emad/redisX/server"
)

func main() {
	server := server.Init("localhost", "6000", "tcp")
	server.RegisterCommand("ping", false, func() string { return "pong123" })
	server.RunAndWait()
}
