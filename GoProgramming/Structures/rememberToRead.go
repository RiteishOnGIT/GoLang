package main

import (
	"fmt"
)

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS

var x int

type person struct {
	first string
	last  string
}

type foo int
var y foo

func main() {
	fmt.Println("Hello, playground")
}
