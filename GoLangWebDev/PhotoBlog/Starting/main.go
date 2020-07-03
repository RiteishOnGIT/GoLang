package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
	c, e := req.Cookie("Cookie")
	if e != nil{
		a, _ := uuid.NewV4()
		c = &http.Cookie{
			Value: a.String(),
		}
		http.SetCookie(w, c)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}
