package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppCongig hold the application config
type AppCongig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProguction  bool
	Session      *scs.SessionManager
}