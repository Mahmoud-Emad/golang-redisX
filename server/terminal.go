package server

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Terminal struct {
	server *RespServer
}

// Init an instance frm the terminal struct to interact with it's methods.
func initTerminal(server *RespServer) Terminal {
	return Terminal{server: server}
}

// TODO: To close using the os module, add the status code as a parameter to modify it.
func (trm Terminal) terminate() {
	os.Exit(1)
}

func (trm Terminal) RaisError(message string, err error) {
	color.Red("\n💣 %s", message)
	color.Red("\n💣 %s", err.Error()+"\n")
	trm.terminate()
}

// Function to print a hello message when starting the server.
func (trm Terminal) Welcome() {
	color.Green("\n🪄  Initializing the server")
	color.White(`
			 _____         __ _    __   __
			|  __ \        | (_)   \ \ / /
			| |__) |___  __| | |___ \ V / 
			|  _  // _ \/ _` + "`" + ` | / __| > <  
			| | \ \  __/ (_| | \__ \/ . \ 
			|_|  \_\___|\__,_|_|___/_/ \_\
	`)
	fmt.Printf(
		"a simple and small %s that can handle the requests that sent from the %s.\n",
		color.GreenString("redis-server"),
		color.GreenString("redis-cli"),
	)
	color.Green("\n🚀 Started RedisX server at %s\n\n", trm.server.port)
}
