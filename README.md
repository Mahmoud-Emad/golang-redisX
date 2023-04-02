<img class="badge" alt="Go report card" tag="github.com/Mahmoud-Emad/golang-redisX" src="https://goreportcard.com/badge/github.com/Mahmoud-Emad/golang-redisX">

<img width="50%" height="50%" alt="Redis logo" src="https://upload.wikimedia.org/wikipedia/en/thumb/6/6b/Redis_Logo.svg/1200px-Redis_Logo.svg.png"/>

# golang-redisX

A minimal functional Redis server written in Golang. I built this to learn Golang while simultaneouly
building out a functional product that begs good code practices, moderate use of concurrent goroutines
and dynamic type management.

## Running locally

Make sure that you have Go installed, and that it supports go modules.

```bash
go run server.go
```

This will open up a mock of `redis-cli`, and you can execute simple basic commands.

```sh
Listening on localhost:6382
redis-cli> GET k
(nil)
redis-cli> SET k 2
OK
redis-cli> GET k
2
redis-cli> GETSET foo bar
bar
redis-cli> 
```

Allowed commands are `GET`, `SET`, `DEL`, `GETSET`, `APPEND`, `SETNX`, `STRLEN`, `SETEX`.

## Running tests

`cd resp && go test`
