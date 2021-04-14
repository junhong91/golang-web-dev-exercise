package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Bark!")
	case "/cat":
		io.WriteString(w, "Kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
