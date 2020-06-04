package  main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	fmt.Println(runtime.NumGoroutine())
	wg.Add(2)
	go func(){
		fmt.Println("HEllo")
		wg.Done()
	}()

	fmt.Println(runtime.NumGoroutine())
	go func(){
		fmt.Println("HEllo1")
		wg.Done()
	}()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	wg.Wait()
	fmt.Println("Hello form anywhere")

}