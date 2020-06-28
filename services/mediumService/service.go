package mediumService

import (
	"fmt"
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/constants"
	"github.com/bozd4g/cherry/proxy/mediumProxy"
	"github.com/bozd4g/cherry/proxy/mediumProxy/mediumProxyDtos"
	"github.com/mitchellh/mapstructure"
	"log"
	"regexp"
	"strings"
)

type mediumService struct {
	MemoryCache caching.IMemoryCache
}

type IMediumService interface {
	GetRss() mediumProxyDtos.RssDto
}

func New(memoryCache caching.IMemoryCache) IMediumService {
	return &mediumService{MemoryCache: memoryCache}
}

func (m *mediumService) GetRss() mediumProxyDtos.RssDto {
	var rssDto mediumProxyDtos.RssDto

	if cache, isExist := m.MemoryCache.Get(constants.RssDataKey); !isExist || cache == nil {
		proxy := mediumProxy.New()
		rss, err := proxy.GetRss()
		if rss == nil || err != nil {
			log.Fatal("An error occured while retrieving to rss!")
			return rssDto
		}

		err = mapstructure.Decode(rss, &rssDto)
		if err != nil {
			log.Fatal("An error occured while decoding to rssDto!")
			return rssDto
		}

		var colNumber = 4
		var rgx = regexp.MustCompile(`(http[s]?:\/\/)?([^\/\s]+\/)(.*)`)

		for i := 0; i < len(rssDto.Items); i++ {
			itemDto := &rssDto.Items[i]
			if len(itemDto.Categories) == 0 {
				rssDto.Items = append(rssDto.Items[:i], rssDto.Items[i+1:]...)
				i--
				continue
			}

			guidMatches := rgx.FindAllStringSubmatch(itemDto.Guid, -1)
			linkMatches := rgx.FindAllStringSubmatch(itemDto.Link, -1)

			itemDto.Id = guidMatches[0][3]
			itemDto.Link = fmt.Sprintf("%s/%s", guidMatches[0][3], strings.Replace(linkMatches[0][3], "/", "-", 10))
			itemDto.ClassName = fmt.Sprintf("col-md-%d", colNumber)

			if i == 0 {
				colNumber = 8
			} else if i == 4 {
				colNumber = 8
			} else {
				colNumber = 4
			}
		}

		m.MemoryCache.Set(constants.RssDataKey, rssDto)
	} else {
		err := mapstructure.Decode(cache, &rssDto)
		if err != nil {
			log.Fatal("An error occured while decoding data to rssDto!")
		}

		log.Println("Data came from cache")
	}

	return rssDto
}
