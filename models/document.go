package models

type IndexDocument struct {
	Title string `json:"title"`
	Rss   Rss    `json:"rss"`
}

type PostDocument struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AboutDocument struct {
	Title string `json:"title"`
}
