
package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// IDataType represents the interface of data types that can be sent 
// over the wire according to RESP.
type IDataType interface {
	isDataType() bool
	ToString() string 
}

// String struct holds a simple non-binary safe string value
type String struct {
	value string
}

func (String) isDataType() bool {
	return true
}

// ToString returns string value
func (s String) ToString() string {
	return s.value
}

// NewString creates new String instance
func NewString(s string) String {
    return String {value: s}
}


func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}