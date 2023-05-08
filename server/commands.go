package server
import "fmt"


type Command struct{
	name string
	withArgs bool
	callableFunction func()
}
