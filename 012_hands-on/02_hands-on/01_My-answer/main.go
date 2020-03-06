package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  region
}

type region struct {
	Sothenrn, Central, Northern bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "Hotel Araguaia",
			Address: "Av. Araguaia",
			City:    "São Félix do Araguaia",
			Zip:     "Do not close",
			Region: region{
				Sothenrn: false,
				Central:  false,
				Northern: true,
			},
		},
		hotel{
			Name:    "Hotel kuriala",
			Address: "Meio do Mato",
			City:    "São Félix do Araguia",
			Zip:     "Don't know what times close",
			Region: region{
				Sothenrn: false,
				Central:  true,
				Northern: false,
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	checkErr(err)

	nf, err := os.Create("index.html")
	checkErr(err)

	err = tpl.Execute(nf, hotels)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
