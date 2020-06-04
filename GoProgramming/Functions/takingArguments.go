package main

import (
	"fmt"
)

func main() {
	foo1()
	bar("James")
}

func foo1() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}