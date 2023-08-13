package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	// Check if template already in cache
	_, inMap := tmplCache[t]
	if !inMap {
		// Create the template
		err = createTemplateCache(t)
		log.Println("New template created")
		if err != nil {
			log.Println(err)
		}
	} else {
		// Template already in cache
		log.Println("Cached template used")
	}

	tmpl = tmplCache[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.html",
	}

	// Parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add template to cache map
	tmplCache[t] = tmpl

	return nil
}
