package main

import "fmt"

//anonymous is said because they do not have any name

func main(){
	person := struct {
		age int
		name string
	}{
		age : 100,
		name : "Ritesh",
	}

	fmt.Println(person)
}