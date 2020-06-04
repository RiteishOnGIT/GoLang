package main

import "fmt"

func main(){
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(*b , b ,&b, *&b)
	fmt.Println(*&a)

	*b = 23123
	fmt.Println(a)
}
