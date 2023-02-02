package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloprint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!! \n")
	switch r.URL.Path {
	case "/hello":
		fmt.Fprint(w, "Hello World")

	case "/hello/luffy":
		fmt.Fprint(w, "Ace ki MKC 2nd time")
	}
}

func helloprintnew(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Timeout function")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Timeout function happening")
}
func ninjamux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Server mux handler</h1>")
}
func main() {

	http.HandleFunc("/", helloprintnew)
	http.HandleFunc("/timeout", timeout)
	//http.HandleFunc("/hello", helloprint)
	//http.ListenAndServe(":5000", nil)

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxserverNinja http.ServeMux
	server.Handler = &muxserverNinja
	muxserverNinja.HandleFunc("/mux", ninjamux)

	server.ListenAndServe()
}
