package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufaturer string
	Model       string
	Doors       int
}

type item struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Harder never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufaturer: "Ford",
		Model:       "F150",
		Doors:       2,
	}

	c := car{
		Manufaturer: "Toyota",
		Model:       "Corolla",
		Doors:       4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.Execute(nf, data)
	checkErr(err)
}

func checkErr(err error) {

}
