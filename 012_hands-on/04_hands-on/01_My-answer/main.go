package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type restaurant struct {
	Name  string
	Local string
	Items []item
}

type item struct {
	Food, Descrip, Meal string
	Price               float64
}

type restaurants []restaurant

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	res := restaurants{
		restaurant{
			Name:  "Casa da Sogra",
			Local: "Vila Santo Antônia",
			Items: []item{
				item{
					Food:    "Oatmeal",
					Descrip: "yum yum",
					Meal:    "Breakfast",
					Price:   4.95,
				},
				item{
					Food:    "Hamburguer",
					Descrip: "Delicous good eating for you",
					Meal:    "Lunch",
					Price:   6.95,
				},
				item{
					Food:    "Pasta Bolognese",
					Descrip: "From Italy delicious eating",
					Meal:    "Dinner",
					Price:   7.95,
				},
			},
		},
		restaurant{
			Name:  "Solar das Águas",
			Local: "Av. Araguaia",
			Items: []item{
				item{
					Food:    "Oatmeal",
					Descrip: "yum yum",
					Meal:    "Brakfast",
					Price:   4.95,
				},
				item{
					Food:    "Hamburguer",
					Descrip: "Delicious good eating fo you",
					Meal:    "Lunch",
					Price:   6.95,
				},
				item{
					Food:    "Pasta Bolognese",
					Descrip: "From Italy delicious eating",
					Meal:    "Dinner",
					Price:   7.95,
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", res)
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", res)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
