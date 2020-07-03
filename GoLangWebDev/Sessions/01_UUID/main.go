package main

import(
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func main(){
	http.HandleFunc("/", session)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func session(w http.ResponseWriter, req *http.Request){
	cookie ,err := req.Cookie("Session")
	if err != nil{
			id, _ := uuid.NewV4()
			cookie := &http.Cookie{
			Name: "Session",
			Value:id.String(),
			Path: "/",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)

}
