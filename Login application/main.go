package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handlefunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	case "/hello":
		fmt.Fprint(w, "hello")
	}

}

var userDb = map[string]string{
	"saif":     "chutiya",
	"pachouri": "bhopal",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	password := r.FormValue("password")
	var newFilename = "logout.html"

	if userDb[username] == "" {
		w.WriteHeader(http.StatusNotFound)
		newFilename = "errorpage.html"
		t, err := template.ParseFiles("Login application/errorpage.html")
		if err != nil {
			fmt.Println("Encountered error", err)
			return
		}
		err = t.ExecuteTemplate(w, newFilename, username)

		if err != nil {
			fmt.Println("Error when executing template", err)
			return
		}
		return
	}
	if userDb[username] == password {
		w.WriteHeader(http.StatusOK)
		t, err := template.ParseFiles("Login application/logout.html")
		if err != nil {
			fmt.Println("Encountered error", err)
			return
		}
		//fmt.Fprintf(w, "Logged in succesfully")
		err = t.ExecuteTemplate(w, newFilename, username)
		if err != nil {
			fmt.Println("Error when executing template", err)
			return
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		newFilename = "errorpage.html"
		t, err := template.ParseFiles("Login application/errorpage.html")
		if err != nil {
			fmt.Println("Encountered error", err)
			return
		}
		err = t.ExecuteTemplate(w, newFilename, username)
		if err != nil {
			fmt.Println("Error when executing template", err)
			return
		}

		//<form action="/login">
		//<input type="submit" name="" id="" value="Go back to login page">
		//</form>
		//fmt.Fprintf(w, "Incorrect credentials")
	}
	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)
}

func login(w http.ResponseWriter, r *http.Request) {
	var filename = "login.html"
	t, err := template.ParseFiles("Login application/login.html")
	if err != nil {
		fmt.Println("Encountered error", err)
		return
	}

	err = t.ExecuteTemplate(w, filename, "Login Page template")
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

func main() {
	http.HandleFunc("/", handlefunc)
	err := http.ListenAndServe("", nil)
	//err := http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)

	if err != nil {

		fmt.Println("Error whlie listening ")
		return
	}
}
