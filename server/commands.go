package server

type Command struct {
	name             string
	withArgs         bool
	callableFunction func()
}
