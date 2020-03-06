package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type item struct {
	Name, Descrip, Meal string
	Price               float64
}

type itens []item

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := itens{
		item{
			Name:    "Oatmeal",
			Descrip: "yum yum",
			Meal:    "Breakfast",
			Price:   4.95,
		},
		item{
			Name:    "Hamburguer",
			Descrip: "Delicious good eating for you",
			Meal:    "Lunch",
			Price:   6.95,
		},
		item{
			Name:    "Pasta Bolognese",
			Descrip: "From Italy delicious eating",
			Meal:    "Dinner",
			Price:   7.95,
		},
	}

	err := tpl.Execute(os.Stdout, m)
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.Execute(nf, m)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
