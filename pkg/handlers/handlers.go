package handlers

import (
	"net/http"

	"github.com/bartosz-bartosz/go-web/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}

func Test(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "test.html")
}
