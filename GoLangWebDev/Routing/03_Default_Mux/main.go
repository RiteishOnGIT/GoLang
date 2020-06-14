package main

import(
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Bingo")
}

func main(){
	http.HandleFunc("/God", d)

	http.ListenAndServe(":8080", nil)
}