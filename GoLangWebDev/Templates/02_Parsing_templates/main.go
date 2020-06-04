package main

import (
	"log"
	"os"
	"text/template"
)

func main(){
	tpl, err := template.ParseFiles("E:\\GoLang\\GoLangWebDev\\Templates 02\\tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	nf, err := os.Create("index.html")
	if err != nil{
		log.Println("Error creating a file", err)
	}
	defer nf.Close()

	err = tpl.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln(err)
	}
}
