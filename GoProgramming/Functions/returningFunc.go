package main

import (
	"fmt"
)

func main() {
	s1 := fooo()
	fmt.Println(s1)

	x := baar()

	fmt.Printf("%T\n", x)

	i := x(34)
	fmt.Println(i)

}

func fooo() string {
	s := "Hello world"
	return s
}

func baar() func(x int) int {
	return func(x int ) int {
		return x
	}
}
