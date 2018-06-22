package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", indexRoute)

	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
