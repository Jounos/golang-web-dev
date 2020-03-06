package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Jesus")
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.ExecuteTemplate(nf, "index.gohtml", "Jesus")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
