
package main

import (
"fmt"
	"runtime"
)

var x int8 = 127

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
