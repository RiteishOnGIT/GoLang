package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type Person struct {
	Name  string
	Motto string
}
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	b := Person{
		Name:  "Mohit",
		Motto: "The belief of no beliefs",
	}

	g := Person{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := Person{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}


	Sages := []Person{b, g, m}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Sages)
	if err != nil {
		log.Fatalln(err)
	}
}