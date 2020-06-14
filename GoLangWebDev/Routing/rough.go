package main


import(
	"log"
	"net/http"
	"text/template"
)

func main(){
	http.Handle("/me", http.HandlerFunc(me))
	http.Handle("/about",http.HandlerFunc(about))
	http.ListenAndServe(":8080", nil)
}

func me(w http.ResponseWriter, r *http.Request){
	tpl, err := template.ParseFiles("somethingnew.gohtml")
	if err != nil{
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "somethingnew.gohtml", "Ritesh")
}

func about(w http.ResponseWriter, r *http.Request){
	tpl, err := template.ParseFiles("somethingnew.gohtml")
	if err != nil{
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "somethingnew.gohtml", "Kaggle")
}