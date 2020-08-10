package mediumClient

type RssDto struct {
	Status string    `json:"status"`
	Feed   FeedDto   `json:"feed"`
	Items  []ItemDto `json:"items"`
}

type FeedDto struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ItemDto struct {
	Id          string   `json:"id"`
	DetailLink  string   `json:"detailLink"`
	Title       string   `json:"title"`
	PubDate     string   `json:"pubDate"`
	Link        string   `json:"link"`
	Guid        string   `json:"guid"`
	Thumbnail   string   `json:"thumbnail"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Categories  []string `json:"categories"`
	ClassName   string   `json:"className"`
}
