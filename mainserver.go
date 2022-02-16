package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\n Hello I am the tcp server")
		fmt.Fprintln(conn, "\nHow's your day?")
		fmt.Fprintf(conn, "%v", "Good day boss!")

		conn.Close()
	}
}
