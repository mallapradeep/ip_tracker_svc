package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy ")
	case "/cat":
		io.WriteString(res, "KITTY, kittty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
