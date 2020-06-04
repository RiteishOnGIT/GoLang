package main

import (
	"fmt"
)

func main() {
	//initialization at the same time
	if x := 42; x == 42 {
		fmt.Println("001")
	}
	//scope of x is only in the if condition above
	// fmt.Println(x) <-- This wont work
	y := 12
	if y == 19{
		fmt.Printf("%d %s",y,"if part")
	} else if y == 12 {
		fmt.Printf("%d %s", y, "else if part")
	} else {
		fmt.Printf("%d %s",y,"else part")
	}
}
