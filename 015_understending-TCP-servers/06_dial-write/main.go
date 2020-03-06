package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	checkErr(err)
	defer conn.Close()

	fmt.Fprintf(conn, "I dialed you.")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
