package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "mow mow mow")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.HandleFunc("/cow", m)

	http.ListenAndServe(":8080", nil)
}
