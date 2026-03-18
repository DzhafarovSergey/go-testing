package main

import (
	"fmt"
	"github.com/DzhafarovSergey/go-testing.git/hello"
)

func main() {
	name := hello.Hello("Andrey")
	fmt.Println(name)
}
