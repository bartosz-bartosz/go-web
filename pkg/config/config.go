package config

import (
	"html/template"
)

// Holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
