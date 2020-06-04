package main

import "fmt"

func main(){
	i := 0
	//This is new to me so remember
	for{
		if(i > 9){
			break
		}
		fmt.Printf("%d\n",i)
		i++;
	}
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}