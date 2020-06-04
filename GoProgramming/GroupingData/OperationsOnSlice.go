package main

import "fmt"

func main(){
	x := []int {1,2,3,4,5}
	fmt.Println(x[0:3])

	for i:= 0 ; i < len(x) ; i++{
		fmt.Println(x[i])
	}
	//Appending Values to the Slice
	y :=[]int {312,4,3243,2534,63,464,57,568,6}
	x = append(x, y...)
	fmt.Println(x)

	//Deleting form a Slice
	x  = append(x[:2],x[4:]...)
	fmt.Println(x)

	//Slice using Make keyword
	z := make([]int, 10, 12)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z[0] = 42
	z[9] = 999

	z = append(z, 3424)

	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 3425,12)

	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
}