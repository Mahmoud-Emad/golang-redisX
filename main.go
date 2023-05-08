package main

import (
	"github.com/Mahmoud-Emad/redisX/server"
	"fmt"
)

func main() {
	// Server Case
	command1 := server.Command{
		name: "test",
		withArgs: false,
		callableFunction: func() { fmt.Println("initial") }
	}
	fmt.Println(command1)
	server := server.New("localhost", "6000", "tcp")
	server.RunAndWait()
}
