package main

import (
	"github.com/Mahmoud-Emad/redisX/server"
)

func main() {
	// Server Case
	server := server.New("localhost", "6000", "tcp")
	server.RunAndWait()
}
