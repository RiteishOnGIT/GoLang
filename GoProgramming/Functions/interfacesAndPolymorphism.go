
package main

import "fmt"

type person1 struct{
	name string
	age int
}
type secretAgent1 struct {
	person1
	ltk bool
}

type human interface {
	speak()
}
// func ( receiver ) identifiers (parameters) (returns) {code return}

func (s secretAgent1) speak(){
	fmt.Println("I am = ", s.name , "with license to kill = " , s.ltk)
}
func (pz person1) speak(){
	fmt.Println("I am = ", pz.name , "with age = " , pz.age)
}

func barr(h human){
	//fmt.Printf("%T", h)
	switch h.(type){
	case person1:
		fmt.Println(h.(person1).name)
	case secretAgent1:
		fmt.Println(h.(secretAgent1).ltk)
	}
	fmt.Println("i was passed into barrr", h)
}
func main() {
	sal1 := secretAgent1{
		person1 : person1{
			name : "raja ",
			age : 12,
		},
		ltk : true,
	}
	p := person1{
		name: "Dr.",
		age:  23,
	}

	sal1.speak()
	barr(p)
	barr(sal1)

}