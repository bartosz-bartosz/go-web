package handlers

import (
	"net/http"

	"github.com/bartosz-bartosz/go-web/pkg/config"
	"github.com/bartosz-bartosz/go-web/pkg/models"
	"github.com/bartosz-bartosz/go-web/pkg/render"
)

// Repository used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html", &models.TemplateData{})
}

// Test page handler
func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, test."
	render.RenderTemplate(w, "test.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
