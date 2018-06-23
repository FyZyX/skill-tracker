package controllers

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Handle GET requests
	case http.MethodGet:
		views.ExecuteTemplate(w, routes.index, nil)
	}
}
