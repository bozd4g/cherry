package models

import (
	"html/template"

	"github.com/bozd4g/cherry/services/mediumService"
)

type IndexDocument struct {
	Title string                  `json:"title"`
	Posts []mediumService.PostDto `json:"posts"`
}

type PostDocument struct {
	Title string        `json:"title"`
	Body  template.HTML `json:"body"`
}

type AboutDocument struct {
	Title string `json:"title"`
}

type ContactDocument struct {
	Title string `json:"title"`
}
