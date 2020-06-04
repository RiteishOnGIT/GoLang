package main

import (
	"fmt"
)

func main(){
	//UNBUFFERED CHANNELS ---------Bidirectional
	c := make(chan int)
	go func(){
		c <- 12
	}()
	fmt.Println(<-c)
	//BUFFERED CHANNELS  ----------Bidirectional
	b := make(chan int, 2)
	b <- 12
	b <- 3
	fmt.Println(<-b, <-b)
}