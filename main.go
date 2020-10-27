package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/JackWinterburn/Go-gmail-client/send"
)

var templates = template.Must(template.ParseFiles(
	"templates/home.html",
	"templates/login.html",
	"templates/404.html",
	"templates/view.html",
	"templates/create-email.html"))

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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "view.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createEmailHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "create-email.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	//serving static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//home page
	http.HandleFunc("/", homeHandler)

	//login page
	http.HandleFunc("/login", loginHandler)

	//view emails
	http.HandleFunc("/view", viewHandler)

	//create emails
	http.HandleFunc("/create-email", createEmailHandler)

	//send emails
	http.HandleFunc("/send", send.CreateAndSendEmail)

	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
