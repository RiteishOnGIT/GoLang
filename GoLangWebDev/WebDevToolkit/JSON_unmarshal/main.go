package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type thumbnail struct {
	URL string
	Height int
	Width int
}

type autogen struct{
	Width int
	Height int
	Title string
	Thumbnail thumbnail
	Animated bool
	IDS []int
}
var tpl *template.Template

func init(){
	tpl, _ = template.ParseFiles("index.gohtml")
}
var data autogen
func main(){
	t := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`
	err := json.Unmarshal([]byte(t), &data)
	if err!= nil{
		log.Fatalln(err)
	}
	fmt.Println(data.IDS[0])
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter,req *http.Request){
	tpl.ExecuteTemplate(w, "index.gohtml",data)
}