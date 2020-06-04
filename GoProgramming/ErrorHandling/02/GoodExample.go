package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("names.txt")
	if err != nil {
		log.Panic(err)
		//log.Fatalln(err)
		//panic(err)
	}

}

func foo(){
	fmt.Print("Defered Functions Dont run","\n")
}

//package main

//import "fmt"

//func main() {
//	b()
//}

//func b() {
//	for i := 1; i < 5; i++ {
//		defer fmt.Println(i)
//	}
//}
