package controllers

import (
	"net/http"
	"skillTracker/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Handle GET requests
	case http.MethodGet:
		data := struct{
			Skill models.Skill
		}{
			models.NewSkill("Single digit addition", 0.31),
		}

		views.ExecuteTemplate(w, routes.index, data)
	}
}
