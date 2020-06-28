package models

import (
	"github.com/bozd4g/cherry/proxy/mediumProxy/mediumProxyDtos"
	"html/template"
)

type IndexDocument struct {
	Title string                 `json:"title"`
	Rss   mediumProxyDtos.RssDto `json:"rss"`
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
