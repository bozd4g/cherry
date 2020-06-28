package utils

import (
	"github.com/bozd4g/cherry/constants"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = make(map[string]*template.Template)

func LoadTemplates() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	layoutTemplate := "templates/layout/_base.html"

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range allTemplates.Templates() {
		isMatch, _ := regexp.MatchString(constants.HtmlRegex, item.Name())
		if !isMatch {
			continue
		}

		templates[item.Name()], err = template.ParseFiles(layoutTemplate, "templates/" + item.Name())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates[tmpl].Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
