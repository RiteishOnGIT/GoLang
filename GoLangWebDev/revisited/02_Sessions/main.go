package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First string
	Last string
}

var tpl *template.Template
var dsessions = map[string]string{}
var dbusers = map[string]user{}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){

	c, err := req.Cookie("session")
	if err != nil{
		sID, _ :=uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	var u1 user

	if un, ok := dsessions[c.Value]; ok{
		u1 = dbusers[un]
	}
	if u1.First != ""{

		http.Redirect(w, req, "/bar", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost{
		u1 = user{
			req.FormValue("username"),
			req.FormValue("firstname"),
			req.FormValue("lastname"),
		}
		dsessions[c.Value] = req.FormValue("username")
		dbusers[req.FormValue("username")] = u1
	}
	fmt.Println(dsessions)
	fmt.Println(dbusers)
	tpl.ExecuteTemplate(w, "index.gohtml", u1)
}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dsessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbusers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
//http.Redirect(w, req, "/", http.StatusSeeOther)
// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV