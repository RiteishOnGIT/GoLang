
package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	default:
		fmt.Println("this is default")
	}
	
}
