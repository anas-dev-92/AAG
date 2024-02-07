package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool // this check to see if all up tp date or still have some changes that never came on the web app
	TemplateCache map[string]*template.Template
	Inproduction  bool
	Session       *scs.SessionManager
}
