package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle string
	Content   template.HTML
}

func main() {
	tmpl := template.Must(template.ParseFiles("base.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Home",
			Content:   template.HTML("Hello World!"),
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":42069", nil)
	if err != nil {
		return
	}
}
