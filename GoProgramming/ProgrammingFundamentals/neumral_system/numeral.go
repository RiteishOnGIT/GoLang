package main

import (
	"fmt"
)

func  main(){
	b := "H"
	bs := []byte(b)
	fmt.Println(b)
	fmt.Println(bs)
	n := bs[0]
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%x\n", n)
}