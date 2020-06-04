package main

import (
	"fmt"
)

type p1 struct {
	age int
	name string
}
type p2 struct {
	health int
	p1
}

func main(){
	person := p2{
		health: 100,
		p1:p1{
			age : 10,
			name: "HariBhai",
		},
	}
	fmt.Println(person)
	fmt.Println(person.name,person.age,person.health)
}
