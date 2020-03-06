package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	checkErr(err)
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("Connection Time Out")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s \n", ln)
	}
	defer conn.Close()

	// new we get hear
	// the connection will time out
	//that breaks us out of scanner loop
	fmt.Println("****COODE GOT HERE****")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
