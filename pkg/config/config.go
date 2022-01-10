package config

import "text/template"

// holds the apps config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
