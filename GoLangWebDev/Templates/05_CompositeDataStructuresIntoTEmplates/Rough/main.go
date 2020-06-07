package main

import (
	"html/template"
	"log"
	"os"
)

var tpl  *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}


type person struct {
	Name string
	LastName []string
}
func main(){
	 p := person{Name: "ritesh",
		 LastName: []string{"jangid", "sharma"},
	 }

	 err := tpl.Execute(os.Stdout, p)
	 if err!= nil{
	 	log.Fatalln(err)
	 }

}
