package main

import (
	"net/http"
	"html/template"
)

type Handler struct{}

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var handler Handler

	http.ListenAndServe(":8080", handler)
}
