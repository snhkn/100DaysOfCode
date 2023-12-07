package handlers

import (
	"net/http"

	"github.com/snhkn/100DaysOfCode/Go/HelloWorld/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
