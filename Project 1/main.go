package main

import (
	"fmt"
	"net/http"
)

func helloprint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!! \n")
	switch r.URL.Path {
	case "/hello":
		fmt.Fprint(w, "Hello World")

	case "/luffy":
		fmt.Fprint(w, "Ace ki MKC")
	}
}

func main() {

	http.HandleFunc("/", helloprint)
	http.ListenAndServe("", nil)
}
