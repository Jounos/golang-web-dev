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

	var str []string = []string{}

	err := tpl.Execute(os.Stdout, str)
	if err != nil {
		log.Fatalln(err)
	}
}
