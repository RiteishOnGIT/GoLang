package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogfile)
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="toby.jpg">`)
}

func dogfile(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "toby.jpg")
}