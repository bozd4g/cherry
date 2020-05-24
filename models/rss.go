package models

type Rss struct {
	Status string `json:"status"`
	Feed   Feed   `json:"feed"`
	Items  []Item `json:"items"`
}

type Feed struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Item struct {
	Title       string   `json:"title"`
	PubDate     string   `json:"pubDate"`
	Link        string   `json:"link"`
	GUID        string   `json:"guid"`
	Author      string   `json:"author"`
	Thumbnail   string   `json:"thumbnail"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Enclosure   struct{} `json:"enclosure"`
	Categories  []string `json:"categories"`
}
