package main


import(
	"fmt"
	"net/http"
)

type pimpo int

func (p pimpo) ServeHTTP(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Ritesh-key", "wohoooooo")
	w.Header().Set("conent-type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "O BHYII HO GYA YEH TOH")

}

func main(){
	var d pimpo
	http.ListenAndServe(":8080", d)
}