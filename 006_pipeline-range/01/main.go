package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, []string{"Gandhi", "MLK", "Budha", "Jesus", "Muhammad"})
	if err != nil {
		log.Fatalln(err)
	}
}
