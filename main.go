package main

import (
	"net/http"
	"html/template"
	"log"
	"os"
	"fmt"
	"io"
)

var bodies []string

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bodies", bodyHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

var indexTemplate = template.Must(template.ParseFiles("template/root.html", "template/index.html"))
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, handler, err := r.FormFile("images")
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("I'm gonna give it to you straight. I don't know what happened there. Try again."))
			return
		}
		defer file.Close()

		if err == nil {
			f, err := os.OpenFile("./static/img/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		}


		w.Write([]byte("success. reload to upload more."))
		return
	}
	renderTemplate(w, indexTemplate, nil)
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	for _, body := range bodies {
		w.Write([]byte(body))
	}
}
func renderTemplate(w http.ResponseWriter, t *template.Template, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}