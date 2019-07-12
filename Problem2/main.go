package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bunty)
	http.HandleFunc("/me", Pradeep)
	//for defaultServerMux we put nil in ListenAndServe
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bunty(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I miss you!")
}

func Pradeep(res http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "something.gohtml", "Pradeep")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
