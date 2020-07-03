package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request){
	fmt.Println("foo method is -> ", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request){
	fmt.Println("bar method is (it will not called next time) -> ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}