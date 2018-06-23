package controllers

import "html/template"

var routes = struct {
	index string
}{
	index: "index.gohtml",
}

var views *template.Template

func init() {
	// Parse templates
	views = template.Must(template.ParseGlob("templates/*.gohtml"))
}
