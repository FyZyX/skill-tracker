package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func init() {
	// Parse templates
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	// Define route handlers
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", indexRoute)

	// Start HTTP server
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Handle GET requests
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	}
}
