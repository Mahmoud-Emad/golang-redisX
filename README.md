<img class="badge" alt="Go report card" tag="github.com/Mahmoud-Emad/golang-redisX" src="https://goreportcard.com/badge/github.com/Mahmoud-Emad/golang-redisX">

<img width="50%" height="50%" alt="Redis logo" src="https://upload.wikimedia.org/wikipedia/en/thumb/6/6b/Redis_Logo.svg/1200px-Redis_Logo.svg.png"/>

# golang-redisX

A minimal functional Redis server written in Golang. I built this to learn Golang while simultaneouly
building out a functional product that begs good code practices, moderate use of concurrent goroutines
and dynamic type management.

## Running locally

Make sure that you have Go installed, and that it supports go modules.

```bash
go run main.go
# the main module added there as code example.
```

This will create and run server on port 6000, you can connect on this server from the `redis-cli` client then interact with it.

To have your own instance and register your custom command with its functions you can do

```go
package main

import (
 "github.com/Mahmoud-Emad/redisX/server"
)

func main() {
  // create a new server on port 6000 
 server := server.Init("localhost", "6000", "tcp")

  // register command with name and function 
 server.RegisterCommand("ping", false, func() string { return "PONG" })
  // in case your function returnes array or number type, just write them as string also like  
 server.RegisterCommand("count", false, func() string { return "123" })
  // There are a few types doesn't implemented yet
  
  // Then run the server and wait for requests  
 server.RunAndWait()
}

```

and from your `redis-cli` client do

```sh
    redis-cli -p 6000
```

This command should connect to our server, now you cn interact with it by using your commands.
