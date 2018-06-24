package controllers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Handle GET requests
	case http.MethodGet:
		data := struct {
			Mastery int
		}{
			25,
		}
		views.ExecuteTemplate(w, routes.index, data)
	}
}
