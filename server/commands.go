package server


type Command struct{
	name string
	argsType []string
	callableFunction func
	
}