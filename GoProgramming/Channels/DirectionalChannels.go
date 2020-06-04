package main

import (
	"fmt"
)

func main(){
	c := make(chan int)
	cs := make(<- chan int) //channel reciever

	cr := make(chan <- int) //channel sender
	fmt.Printf("%T\n",cs)
	fmt.Printf("%T\n",cr)
	fmt.Printf("%T\n",c)

	// general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
	//fmt.Println(<-c)
}
