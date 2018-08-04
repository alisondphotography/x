package main

import (
	"net/http"
	"html/template"
	"log"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal("8080", nil)
}

var indexTemplate = template.Must(template.ParseFiles("template/root.html", "template/index.html"))
func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, indexTemplate, nil)
}
func renderTemplate(w http.ResponseWriter, t *template.Template, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}