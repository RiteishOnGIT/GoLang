package main

import (
	"encoding/json"
	"fmt"
	)

type person struct {
	FirstName string
	LastName string
	Age int
}


func main(){
	p1 := person{
		FirstName: "James",
		LastName: "bond",
		Age: 23,
	}

	p2 := person{
		FirstName: "Miss",
		LastName: "MoneyPenny",
		Age: 32,
	}

	people := []person{p1,p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(bs)
		fmt.Println(string(bs))
	}
}
