package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bartosz-bartosz/go-web/pkg/config"
)

var app *config.AppConfig

// Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// Get the template cache from app config
	templateCache := app.TemplateCache

	// Get template from cache
	template, ok := templateCache[tmpl]
	log.Println(template)
	if !ok {
		log.Fatal("Template not found:", tmpl)
	}

	buf := new(bytes.Buffer)

	_ = template.Execute(buf, nil)

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
		log.Println(myCache)
	}

	return myCache, nil
}
