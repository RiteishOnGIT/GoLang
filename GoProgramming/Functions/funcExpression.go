package main

import "fmt"

func main(){

	f := func(x int ) int{
		return x
	}(23)
	fmt.Println(f)
}
