package models

type IndexDocument struct {
	Title string `json:"title"`
}

type PostDocument struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AboutDocument struct {
	Title string `json:"title"`
}
