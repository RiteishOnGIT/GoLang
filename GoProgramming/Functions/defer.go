package main

import (
	"fmt"
)

func main() {
	fmt.Println("THIS")
	defer foo4()
	bar3()
}

func foo4() {
	fmt.Println("foo")
}

func bar3() {
	fmt.Println("bar")
}
