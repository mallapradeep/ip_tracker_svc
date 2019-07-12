package main

import (
	"io"
	"log"
	"net"
)

//writing to the connection
func main() {
	//Listen
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		//Accept
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		//write to connection
		io.WriteString(conn, "I see you connected...")

		conn.Close()
	}
}
