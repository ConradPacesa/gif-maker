package config

import "html/template"

// TPL points to the html tenplates
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.html"))
}
