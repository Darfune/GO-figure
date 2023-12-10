package handlers

import (
	"net/http"
	"web_01/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request)  {
	render.RenderTemplate(w, "about.page.tmpl")
}


