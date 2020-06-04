package main

import (
	"fmt"
)

func main() {
	fo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}

func fo() {
	fmt.Println("foo ran")
}