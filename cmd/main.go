package main

import (
	"fmt"
	"github.com/DzhafarovSergey/go-testing.git/even"
	"github.com/DzhafarovSergey/go-testing.git/hello"
)

func main() {
	name := hello.Hello("Andrey")
	fmt.Println(name)

	number_even := even.IsEven(4)
	fmt.Println(number_even)

	number_no_even := even.IsEven(5)
	fmt.Println(number_no_even)
}
