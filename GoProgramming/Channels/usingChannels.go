package main

import "fmt"

func main(){
	c := make(chan int)
	//send

	go foo(c)
	//recieve
	bar(c)

	fmt.Println("About to exit")
}

func foo(c chan <- int){
	c <- 12
}

func bar (c <-chan int){
	fmt.Println(<-c)
}