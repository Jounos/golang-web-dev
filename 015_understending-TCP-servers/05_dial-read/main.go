package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	checkErr(err)
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	checkErr(err)

	fmt.Println(string(bs))
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
