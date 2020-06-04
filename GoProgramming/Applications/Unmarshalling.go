package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main(){
	s :=`[{"FirstName":"James","LastName":"bond","Age":23},{"FirstName":"Miss","LastName":"MoneyPenny","Age":32}]`
	s1 := []byte(s)

	//people := []person{}
	var people []person

	err := json.Unmarshal(s1, &people)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("All of the data : " ,people)
	for i := range people{
		fmt.Println(people[i].FirstName,people[i].LastName,people[i].Age)
	}
}