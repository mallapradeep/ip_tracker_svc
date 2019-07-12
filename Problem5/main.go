package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

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

		fmt.Println("Code got here")
		io.WriteString(conn, "I see you connected")i

		conn.Close()
	}
}
