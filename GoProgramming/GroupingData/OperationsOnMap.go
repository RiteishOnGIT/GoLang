package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	//inserting elements to the map

	m["ritesh"] = 20
	m["ritesh"] = 20
	m["ritesh"] = 20
	for i,v := range m{
		fmt.Println(i ,v)
	}
	delete(m,"ritesh")

	for i,v := range m{
		fmt.Println(i ,v)
	}
}
