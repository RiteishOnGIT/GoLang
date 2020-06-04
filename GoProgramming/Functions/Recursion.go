package main

import "fmt"

func main()  {
	x := Factorial(5)
	fmt.Println(x)

}

func Factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}