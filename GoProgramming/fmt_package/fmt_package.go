package main

import (
"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)//prints type
	fmt.Printf("%b\n", y)//prints binary
	fmt.Printf("%x\n", y)//prints hex
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v", y)
}

