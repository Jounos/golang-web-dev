package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := restaurants{
		restaurant{
			Name: "Federico's",
			Menu: menu{
				meal{
					Meal: "Brakfast",
					Item: []item{
						item{
							Name:    "Oatmeal",
							Descrip: "yum yum",
							Price:   4.95,
						},
						item{
							Name:    "Cheerios",
							Descrip: "American eating food traditional now",
							Price:   3.95,
						},
						item{
							Name:    "Juice Orange",
							Descrip: "Delicious drinking in throat squeezed fresh",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:    "Hamburguer",
							Descrip: "Delicious good eating for you",
							Price:   6.95,
						},
						item{
							Name:    "Cheese Mested Sandwich",
							Descrip: "Make cheese bread melt grease hot",
							Price:   3.95,
						},
						item{
							Name:    "French Fries",
							Descrip: "French eat potatoe fingers",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:    "Pasta Bolognese",
							Descrip: "From Italy delicious eating",
							Price:   7.95,
						},
						item{
							Name:    "Steak",
							Descrip: "Dead cow dgrilled bloody",
							Price:   13.95,
						},
						item{
							Name:    "Bistro Potatoe",
							Descrip: "Bistro bar wood American bacon",
							Price:   6.95,
						},
					},
				},
			},
		},
		restaurant{
			Name: "Domenicos",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:    "Oatmeal",
							Descrip: "yum yum",
							Price:   4.95,
						},
						item{
							Name:    "Cheerios",
							Descrip: "American eating food traditional now",
							Price:   3.95,
						},
						item{
							Name:    "Juice Orange",
							Descrip: "Delicious drinking inthrout squeezed fresh",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:    "Hamburger",
							Descrip: "Delicious good eating for you",
							Price:   6.95,
						},
						item{
							Name:    "Cheese Melted Sandwich",
							Descrip: "Make Cheese bread melt grease hot",
							Price:   2.95,
						},
						item{
							Name:    "French Fries",
							Descrip: "french eat potatoe fingers",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:    "Pasta Bolognese",
							Descrip: "From Italy delicious eating",
							Price:   7.95,
						},
						item{
							Name:    "Steak",
							Descrip: "Dead cow grilled bloody",
							Price:   13.95,
						},
						item{
							Name:    "bistro Potatoe",
							Descrip: "Bistro bar wood American bacon",
							Price:   6.95,
						},
					},
				},
			},
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
