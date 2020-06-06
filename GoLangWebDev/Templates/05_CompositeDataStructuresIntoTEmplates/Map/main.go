package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){

	v := map[string]string{
		"Bingo" : "ogniB",
		"Lingo" : "ogniL",
		"Pingoo": "oogniP",
	}

	err := tpl.Execute(os.Stdout, v)
	if err != nil{
		log.Fatalln(err)
	}
}