package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", indexRoute)

	http.ListenAndServe(":8080", nil)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
