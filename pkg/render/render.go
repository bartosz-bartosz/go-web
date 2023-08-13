package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bartosz-bartosz/go-web/pkg/config"
	"github.com/bartosz-bartosz/go-web/pkg/models"
)

var app *config.AppConfig

// Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// Get template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Template not found:", tmpl)
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = template.Execute(buf, td)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all HTML files from /templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	// Loop through HTML files
	for _, page := range pages {
		fileName := filepath.Base(page)

		// log.Println("Adding template to cache:", fileName)
		templateSet, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[fileName] = templateSet
	}

	return myCache, nil
}
