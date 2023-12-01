package handler

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
}

func serveWeb(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func serveWeb2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	data := Page{
		Title: "Hello World",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
