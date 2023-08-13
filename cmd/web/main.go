package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bartosz-bartosz/go-web/pkg/config"
	"github.com/bartosz-bartosz/go-web/pkg/handlers"
	"github.com/bartosz-bartosz/go-web/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create templateCache", err)
	}

	app.TemplateCache = templateCache
	render.NewTemplates(&app)

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/test", handlers.Test)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
