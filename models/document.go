package models

type IndexDocument struct {
	Title string `json:"title"`
	Rss   RssDto `json:"rss"`
}

type PostDocument struct {
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	SelectedItem ItemDto `json:"selectedItem"`
}

type AboutDocument struct {
	Title string `json:"title"`
}

type ContactDocument struct {
	Title string `json:"title"`
}
