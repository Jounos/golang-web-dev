package main

import "fmt"

func main() {
	name :="Geovanny N. Liberato"

	tpl := `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset='UTF-8'/>
				<title>Hello World!</titile>
			</head>
			<body>
				<h1>` + name +`</h1>
			</body>
		</html>`
	fmt.Println(tpl)
}
