package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	checkErr(err)

	defer li.Close()

	for {
		conn, err := li.Accept()
		checkErr(err)

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "how is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
