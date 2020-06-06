package main

import (
	"log"
	"os"
	"text/template"
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
	Rit := sage{
		"Ritesh",
		"jangid",
	}
	pit := sage{
		"Pitesh",
		"jangid",
	}
	jit := sage{
		"Jitesh",
		"jangid",
	}
	kit := sage{
		"Kitesh",
		"jangid",
	}
	fit := sage{
		"Fitesh LOL",
		"jangid",
	}

	v := []sage{Rit,pit,jit,kit,fit}
	err := tpl.Execute(os.Stdout, v)
	if err != nil{
		log.Fatalln(err)
	}
}

