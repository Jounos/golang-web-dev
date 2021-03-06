package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
		<QDOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8"/>
				<title>Hello World</title>
			</head>
			<body>
				<h1>` + name + `</h1>
			</body>
		</html>`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error in creating file!", err)
	}
	io.Copy(nf, strings.NewReader(str))
}
