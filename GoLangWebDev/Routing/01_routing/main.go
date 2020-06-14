package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/about":
		io.WriteString(w, "about page")
	case "/home":
		io.WriteString(w, "Home page")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}