package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplates insere os templates html na variável templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecuteTemplate rendeiza uma página html na tela
func ExecuteTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
