package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//listen --> Accept --> Handle Connection -Read or Write

func main() {
	//listen
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	//Accept n make a connection
	for {
		conn, err := li.Accept()
		//if there is error do this
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		//else
		go handle(conn)
	}
}

//handling connection
func handle(conn net.Conn) {
	defer conn.Close()

	//read request
	request(conn)

	//write response
	respond(conn)
}

//Write func for request and response

func request(conn net.Conn) {
	i := 0
	//use bufio fo makin a new connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//request line
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
