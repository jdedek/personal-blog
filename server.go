package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

func renderPage(w http.ResponseWriter, tmplFile string, pageTitle string) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", tmplFile))
	data := Page{
		Title: pageTitle,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "templates/home.html", "Home")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "templates/articles.html", "Articles")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "templates/contact.html", "Contact")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/tech/", ArticlesHandler)
	http.HandleFunc("/contact/", ContactHandler)

	// Serve static assets from the "assets" directory
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":42069", nil)
	if err != nil {
		return
	}
}
