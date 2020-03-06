package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

// Record : struct reponsable to armazenate the data
type Record struct {
	Date string
	Open float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	//parse file
	records := prs("table.csv")

	//parse templete
	tpl, err := template.ParseFiles("hw.gohtml")
	checkErr(err)

	// execute template
	err = tpl.Execute(res, records)
	checkErr(err)
}

func prs(filePath string) (records []Record) {
	src, err := os.Open(filePath)
	checkErr(err)
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	checkErr(err)

	records = make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date.Format("2006-01-02"),
			Open: open,
		})
	}

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
