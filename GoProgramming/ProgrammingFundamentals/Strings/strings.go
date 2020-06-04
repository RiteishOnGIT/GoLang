
package main

import (
"fmt"
)

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}
}
