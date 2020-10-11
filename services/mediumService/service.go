package mediumService

import (
	"fmt"
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/clients/mediumClient"
	"github.com/bozd4g/cherry/constants"
	"github.com/mitchellh/mapstructure"
	"log"
)

type mediumService struct {
	MemoryCache caching.IMemoryCache
}

type IMediumService interface {
	GetPosts() []PostDto
	ClearCache()
}

func New(memoryCache caching.IMemoryCache) IMediumService {
	return &mediumService{MemoryCache: memoryCache}
}

func (m *mediumService) GetPosts() []PostDto {
	var postsDto []PostDto
	if cache, isExist := m.MemoryCache.Get(constants.PostsCacheKey); !isExist || cache == nil {
		mediumClient := mediumClient.New()
		rssDto, err := mediumClient.GetRss()
		if rssDto == nil || err != nil {
			log.Fatal("An error occured while retrieving to rss!")
			return postsDto
		}

		var colNumber = 4

		for i, v := range rssDto.Channel.Item {
			if len(v.Category) == 0 {
				continue
			}

			if i == 0 {
				colNumber = 8
			} else if i == 4 {
				colNumber = 8
			} else {
				colNumber = 4
			}

			post := PostDto{}.Create(v)
			post.ClassName = fmt.Sprintf("col-md-%d", colNumber)

			postsDto = append(postsDto, post)
		}

		m.MemoryCache.Set(constants.PostsCacheKey, postsDto)
	} else {
		err := mapstructure.Decode(cache, &postsDto)
		if err != nil {
			log.Fatal("An error occured while decoding data to postDto!")
		}

		log.Println("Data came from cache")
	}

	return postsDto
}

func (m *mediumService) ClearCache() {
	m.MemoryCache.Flush()
}