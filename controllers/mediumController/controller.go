package mediumController

import (
	"github.com/bozd4g/cherry/caching"
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
	controller.Engine.GET("/api/medium/flush", controller.flush)
}

func (controller *mediumController) feedHandler(c *gin.Context) {
	posts := controller.MediumService.GetPosts()
	c.JSON(http.StatusOK, posts)
}

func (controller *mediumController) flush(c *gin.Context) {
	controller.MediumService.ClearCache()
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

