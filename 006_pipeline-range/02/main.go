package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var peeps []string = []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(nf, peeps)
	if err != nil {
		log.Fatalln(err)
	}
}
