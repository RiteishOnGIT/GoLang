package main

import (
"fmt"
)

func main() {
	y := []int {1,2,3,4,5,6,7}
	x := sum(y...)
	fmt.Println("The total is", x)
}

func sum(x...(int)) int {
	fmt.Printf("%T", x)
	sum := 0
	for i := range x{
		sum = sum + x[i]
	}
	return sum
}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
