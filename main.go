package main

import (
	"github.com/Mahmoud-Emad/redisX/server"
)

func main() {
	// Server Case
	// command1 := server.Command{}
	// command1.name = ""
	// // command1.name
	// // name:             "test",
	// // callableFunction: func() { fmt.Println("initial") },
	// // withArgs:         false,
	// // }
	// fmt.Println(command1)
	server := server.New("localhost", "6000", "tcp")
	server.RunAndWait()
}
