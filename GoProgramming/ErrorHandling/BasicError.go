package main

import (
	"fmt"
)

func main(){
	i,err := fmt.Println(2321)
	if err != nil {
		fmt.Println("Error occured")
	}
	fmt.Println(i)
}
