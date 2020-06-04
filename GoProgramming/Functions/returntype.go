package main

import "fmt"

func main(){

	foo2()
	bar1("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	}

func foo2() {
	fmt.Println("hello from foo")
}

func bar1(s string){
	fmt.Println(s)
}

func woo (s string) string{
	return fmt.Sprint("Hello miss ",s);
}
