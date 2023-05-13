package server

type Command struct {
	Name             string
	WithArgs         bool
	CallableFunction func() string
}
