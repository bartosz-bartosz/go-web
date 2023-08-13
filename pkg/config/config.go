package config

import (
	"html/template"
	"log"
)

// Holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
