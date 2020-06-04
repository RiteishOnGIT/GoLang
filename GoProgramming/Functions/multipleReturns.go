package main

import "fmt"

func main(){

	foo3()
	bar2("James")
	x, y := mouse("Ian fleming", true)
	fmt.Println(x)
	fmt.Println(y)
}

func foo3() {
	fmt.Println("hello from foo")
}

func bar2(s string){
	fmt.Println(s)
}

func mouse (s string, s1 bool) (string,bool){
	s = fmt.Sprint("Hello from ", s)
	s1 = true
	return s,s1
}

