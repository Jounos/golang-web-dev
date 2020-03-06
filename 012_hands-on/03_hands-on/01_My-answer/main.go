package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type restaurant struct {
	Name      string
	Breakfast []menu
	Lunch     []menu
	Dinner    []menu
}

type menu struct {
	Plate string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	res := restaurant{
		Name: "Casa da Sogra",
		Breakfast: []menu{
			menu{
				Plate: "Pão com Ovo",
			},
			menu{
				Plate: "Ovo mexido",
			},
			menu{
				Plate: "Presunto, Queijo e Salsicha",
			},
		},
		Lunch: []menu{
			menu{
				Plate: "Salada",
			},
			menu{
				Plate: "Churrasco",
			},
			menu{
				Plate: "Lasanha",
			},
		},
		Dinner: []menu{
			menu{
				Plate: "Pizza",
			},
			menu{
				Plate: "Hamburguer",
			},
			menu{
				Plate: "Sushi",
			},
		},
	}

	err := tpl.Execute(os.Stdout, res)
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.Execute(nf, res)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
