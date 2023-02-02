package main

import "net/http"

func handlefunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
	}
}

func url(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", handlefunc)
	http.ListenAndServe("", nil)
}
