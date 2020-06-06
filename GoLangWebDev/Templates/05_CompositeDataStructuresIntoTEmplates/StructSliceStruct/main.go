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

type collection struct {
	Sliceinstruct []sage
}

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


	V := []sage{Rit,pit,jit,kit,fit}

	x :=collection{
		V,
	}

	err := tpl.Execute(os.Stdout, x)
	if err != nil{
		log.Fatalln(err)
	}
}

