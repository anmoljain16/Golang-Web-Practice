package main

import (
	"fmt"
	"net/http"
)

func helloprint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func main() {

	http.HandleFunc("/", helloprint)
	http.ListenAndServe("", nil)
}
