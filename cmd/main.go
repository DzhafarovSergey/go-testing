package main

import (
	"fmt"
	"github.com/DzhafarovSergey/go-testing.git/even"
	"github.com/DzhafarovSergey/go-testing.git/hello"
	"github.com/DzhafarovSergey/go-testing.git/reverse"
)

func main() {
	name, err := hello.Hello("Andrey")
	if err != nil {
		fmt.Println(fmt.Sprintf("Hello failed: %v", err))
	}
	fmt.Println(name)

	numberEven := even.IsEven(4)
	fmt.Println(numberEven)

	numberNoEven := even.IsEven(5)
	fmt.Println(numberNoEven)

	reverseHello := reverse.Reverse("hello")
	fmt.Println(reverseHello)
	reverseHead := reverse.Reverse("главный")
	fmt.Println(reverseHead)
	reverseRune := reverse.Reverse("😉🙂")
	fmt.Println(reverseRune)
	reverseEmpty := reverse.Reverse("")
	fmt.Println(reverseEmpty)
}
