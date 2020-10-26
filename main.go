package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"templates/home.html", "templates/login.html", "templates/404.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//stop people from accessing non-existant urls
	if r.URL.Path != "/" {
		_ = templates.ExecuteTemplate(w, "404.html", nil)
		return
	}

	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//home page
	http.HandleFunc("/", homeHandler)

	//login page
	http.HandleFunc("/login", loginHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
	fmt.Println("Server running on port 3000")
}
