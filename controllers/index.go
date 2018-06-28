package controllers

import (
	"net/http"
	"skillTracker/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Handle GET requests
	case http.MethodGet:
		skills := []models.Skill{
			models.NewSkill("Single digit addition", 0.31),
			models.NewSkill("Double digit addition", 0.43),
			models.NewSkill("Single digit subtraction", 0.56),
			models.NewSkill("Single digit multiplication", 0.27),
		}
		data := struct{
			Skills []models.Skill
		}{
			skills,
		}

		views.ExecuteTemplate(w, routes.index, data)
	}
}
