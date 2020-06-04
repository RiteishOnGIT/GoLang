package main

import "fmt"

func main(){
	x := 0
	fmt.Println("X befor passing into func ", x)
	fmt.Println("X's Address ", &x)
	foo(&x)
	fmt.Println("X After passing into func ", x)

}

func foo ( y *int){
	fmt.Println("Y is same as x", *y)
	fmt.Println("Address of y ", &y)
	* y = 200
}
