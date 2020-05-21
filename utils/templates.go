package utils

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = make(map[string]*template.Template)

func LoadTemplates() {
	allTemplates, _ := template.ParseGlob("templates/*.html")
	for _, item := range allTemplates.Templates() {
		isMatch, _ := regexp.MatchString("([a-zA-Z0-9\\s_\\\\.\\-():])+(.html)$", item.Name())
		if !isMatch {
			continue
		}

		templates[item.Name()] = template.Must(template.ParseFiles("templates/layout/_base.html", "templates/" + item.Name()))
	}
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates[tmpl].Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
