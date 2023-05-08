package server

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Terminal struct {
	server *RespServer
}

func initTerminal(server *RespServer) Terminal {
	return Terminal{server: server}
}

func (trm Terminal) Welcome() {
	// Function to print a hello message when starting the server.
	color.White("\n🪄 : Initializing the server")
	fmt.Println(strings.Repeat("-", 75))
	color.Cyan(`
	_____          _ _    __   __
	|  __ \        | (_)   \ \ / /
	| |__) |___  __| |_ ___ \ V / 
	|  _  // _ \/ _` + "`" + ` | / __| > <  
	| | \ \  __/ (_| | \__ \/ . \ 
	|_|  \_\___|\__,_|_|___/_/ \_\
	`)
	fmt.Println("a simple and small redis-server that can handle the requests that sent from the redis-cli.")
	fmt.Println(strings.Repeat("-", 75))
	color.Blue("\n🚀 : Started RedisX server at %s", trm.server.port)
}
