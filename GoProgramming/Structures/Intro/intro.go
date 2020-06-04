package main

import (
	"fmt"
)

type person struct {
	age int
	first string
	last string
}

func main(){
	p1:= person{}
	fmt.Println(p1)
	p1.age =10
	p1.first = "Last"
	p1.last = "First"

	fmt.Println(p1)
}
