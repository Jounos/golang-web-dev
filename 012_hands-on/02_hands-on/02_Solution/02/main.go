package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := region{
		Region: "Southern",
		Hotels: []hotel{
			hotel{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			hotel{
				Name:    "h",
				Address: "4",
				City:    "L",
				Zip:     "95612",
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.Execute(nf, h)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
