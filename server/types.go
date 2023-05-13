package server

// Redis resp server struct
type RespServer struct {
	id       string
	host     string
	port     string
	network  string
	Commands []Command
	client   *ClientConnection
}

type ClientConnection struct {
	id            string
	remoteAddress string
}
