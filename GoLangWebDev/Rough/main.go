package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	FirstName string
}

var USessions = map [string]string{}
var UIDs = map [string]user{}

var tpl *template.Template
func init(){
	tpl = template.Must(tpl.ParseFiles("index.gohtml"))
}

func main(){
	http.HandleFunc("/session",index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request){
	c, err := req.Cookie("session")
	if err!= nil{
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)
	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("UserName")
		f := req.FormValue("FirstName")
		USessions[c.Value] = un
		u = user{un ,f}
		UIDs[un] = u
	}
	fmt.Println(USessions)
	fmt.Println(UIDs)
	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

