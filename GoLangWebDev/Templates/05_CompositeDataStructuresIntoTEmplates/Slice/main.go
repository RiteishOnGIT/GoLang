package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init (){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	variable1 := []string{"bhuda", "mod"}

	err := tpl.Execute(os.Stdout, variable1)
	if err != nil{
		log.Fatalln(err)
	}
}