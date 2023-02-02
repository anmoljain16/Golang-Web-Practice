package main

import "net/http"

func handlefunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)

	case "/login-submit":
		loginSubmit(w, r)
	}

}

func loginSubmit(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", handlefunc)
	http.ListenAndServe("", nil)
}
