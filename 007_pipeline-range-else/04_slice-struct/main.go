package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceaces with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil",
	}

	var sages []sage = []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.Execute(nf, sages)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
