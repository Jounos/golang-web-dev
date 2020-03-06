package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Geovanny N. Liberato"

	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8"/>
				<title>Hello World!</title>
			</head>
			<body>
				<h1>` + name + `</h1>
			</body>
		</html>`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Faild in create file!", err)
	}
	io.Copy(nf, strings.NewReader(str))
}
