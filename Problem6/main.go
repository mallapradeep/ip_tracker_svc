package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

//Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

//Pass the connection of type net.Conn as an argument into this function.

//Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

func main() {
	//listen
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	//Accept
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	//reading frm the connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			//when ln is empty, header is done
			fmt.Println("This is the end of the http request headers")
			break
		}
	}
	io.WriteString(conn, "Here we write to the response.")
}
