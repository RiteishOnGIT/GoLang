package main

import "fmt"

func main(){
	i := 0
	for{
		if i > 100 {
			break
		}
		if i%2 != 0 {
			i++
			continue
		}
		fmt.Printf("%d %s",i," ")
		i++
	}
}

