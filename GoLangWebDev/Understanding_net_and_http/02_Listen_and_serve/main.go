package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this function to write")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}