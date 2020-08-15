package mediumService

import (
	"fmt"
	"github.com/bozd4g/cherry/clients/mediumClient"
	"github.com/bozd4g/cherry/constants"
	"regexp"
	"strings"
)

type PostDto struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	PublishDate string   `json:"publishDate"`
	Link        string   `json:"link"`
	Guid        string   `json:"guid"`
	Author      string   `json:"author"`
	Thumbnail   string   `json:"thumbnail"`
	Content     string   `json:"content"`
	Categories  []string `json:"categories"`
	ClassName   string   `json:"className"`
}

func (p PostDto) Create(itemDto mediumClient.RssItemDto) PostDto {
	var urlRgx = regexp.MustCompile(constants.UrlRegex)
	var cdnRgx = regexp.MustCompile(constants.MediumCdnRegex)

	guidMatches := urlRgx.FindAllStringSubmatch(itemDto.Guid.Text, -1)
	urlMatches := urlRgx.FindAllStringSubmatch(itemDto.Link, -1)

	cdnMatches := cdnRgx.FindAllStringSubmatch(itemDto.Encoded, -1)

	return PostDto{
		Id:          strings.ReplaceAll(guidMatches[0][3], "p/", ""),
		Title:       itemDto.Title,
		PublishDate: itemDto.PubDate,
		Guid:        itemDto.Guid.Text,
		Author:      itemDto.Creator,
		Content:     itemDto.Encoded,
		Categories:  itemDto.Category,
		Thumbnail:   cdnMatches[0][0],
		Link:        fmt.Sprintf("%s/%s", guidMatches[0][3], strings.Replace(urlMatches[0][3], "/", "-", 10)),
	}
}
