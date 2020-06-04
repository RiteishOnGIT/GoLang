package main

import "fmt"

func main(){
	x := []int{1,2,3,4,5}
	z := mysum(x...)
	fmt.Println(z)
}

func mysum(x ...int) int {
	sum := 0
	for i:=0 ;i< len(x); i++{
		sum = sum + x[i]
	}
	return sum
}