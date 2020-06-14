package main

import (
	"io"
	"net/http"
)
type hotdog int
type pimpoo int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "Garam kutta")
	io.WriteString(res, "Grama kutta")
}
func (c pimpoo) ServeHTTP(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "pimpoo billi")
	io.WriteString(res, "pimpoo billi")
}


func main() {
	var d hotdog
	var c pimpoo
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}