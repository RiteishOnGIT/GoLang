package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	x := person{

		 req.FormValue("first"),
		req.FormValue("last"),
		req.FormValue("subscribe") == "on",
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", x)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
