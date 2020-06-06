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

type sage struct {
	Name string
	Last string
};

func main(){
	v := sage{
		"Ritesh",
		"jangid",
	}

	err := tpl.Execute(os.Stdout, v)
	if err != nil{
		log.Fatalln(err)
	}
}

