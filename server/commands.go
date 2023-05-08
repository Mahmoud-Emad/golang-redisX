package server

type Command struct {
	name             string
	withArgs         bool
	callableFunction func()
}

func initCommand() Command {
	return Command{
		name:             "",
		withArgs:         false,
		callableFunction: nil,
	}
}
