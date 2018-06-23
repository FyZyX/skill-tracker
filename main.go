package main

import (
	"net/http"
	"log"

	"skillTracker/controllers"
)

func main() {
	// Define route handlers
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", controllers.IndexHandler)

	// Start HTTP server
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
