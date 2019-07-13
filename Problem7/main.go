package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)
//For the default route "/" Have a func called "foo" which writes to the response "foo ran"
//For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", bunty)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func bunty(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
