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

//open gitbash do-> go run mainserver.go
//open cmd
//type-> telnet localhost 8080
//(8080 or any other port mentioned in the program)
//
