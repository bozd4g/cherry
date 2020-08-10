package mediumController

import (
	"encoding/json"
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/clients/mediumClient"
	"github.com/bozd4g/cherry/services/mediumService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type mediumController struct {
	Engine        *gin.Engine
	MediumService mediumService.IMediumService
}

type IMediumController interface {
	Init()
}

func New(engine *gin.Engine, memoryCache caching.IMemoryCache) IMediumController {
	mediumService := mediumService.New(memoryCache)
	return &mediumController{Engine: engine, MediumService: mediumService}
}

func (controller *mediumController) Init() {
	controller.Engine.GET("/api/medium/feed", controller.feedHandler)
}

func (controller *mediumController) feedHandler(c *gin.Context) {
	rss := controller.MediumService.GetRss()
	response, err := json.Marshal(rss)
	if err != nil {
		emptyResponse, _ := json.Marshal(mediumClient.RssDto{
			Status: "",
			Feed:   mediumClient.FeedDto{},
			Items:  []mediumClient.ItemDto{},
		})
		response = emptyResponse
	}
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "An error occured while retrieving the feed from medium.com"})
		return
	}

	c.JSON(http.StatusOK, response)
}
