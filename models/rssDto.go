package models

type RssDto struct {
	Status string    `json:"status"`
	Feed   Feed      `json:"feed"`
	Items  []ItemDto `json:"items"`
}

type ItemDto struct {
	Id            string   `json:"id"`
	DetailLink    string   `json:"detailLink"`
	Title         string   `json:"title"`
	PubDate       string   `json:"pubDate"`
	Link          string   `json:"link"`
	Guid          string   `json:"guid"`
	Thumbnail     string   `json:"thumbnail"`
	Description   string   `json:"description"`
	Content       string   `json:"content"`
	Categories    []string `json:"categories"`
	ClassName     string   `json:"className"`
}
