package pageController

import (
	"fmt"
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/clients/mediumClient"
	"github.com/bozd4g/cherry/constants"
	"github.com/bozd4g/cherry/models"
	"github.com/bozd4g/cherry/services/mediumService"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type pageController struct {
	Engine        *gin.Engine
	MediumService mediumService.IMediumService
}

type IPageController interface {
	Init()
}

func New(engine *gin.Engine, memoryCache caching.IMemoryCache) IPageController {
	mediumService := mediumService.New(memoryCache)
	return &pageController{Engine: engine, MediumService: mediumService}
}

func (controller *pageController) Init() {
	controller.Engine.GET("/", controller.indexHandler)
	controller.Engine.GET("/p/:id/:title", controller.postHandler)
	controller.Engine.GET("/about", controller.aboutHandler)
	controller.Engine.GET("/contact", controller.contactHandler)
}

func (controller *pageController) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "/index", models.IndexDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "Home"),
		Rss:   controller.MediumService.GetRss(),
	})
}

func (controller *pageController) postHandler(c *gin.Context) {
	rss := controller.MediumService.GetRss()
	var selectedRss mediumClient.ItemDto
	for _, v := range rss.Items {
		if v.Id == fmt.Sprintf("p/%s", c.Param("id")) {
			selectedRss = v
		}
	}

	c.HTML(http.StatusOK, "/post", models.PostDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, selectedRss.Title),
		Body:  template.HTML(selectedRss.Content),
	})
}

func (controller *pageController) aboutHandler(c *gin.Context) {
	//utils.ExecuteTemplate(c, "about.html", models.AboutDocument{
	//	Title: fmt.Sprintf(constants.DocumentTitle, "About"),
	//})
}

func (controller *pageController) contactHandler(c *gin.Context) {
	//utils.ExecuteTemplate(c, "contact.html", models.ContactDocument{
	//	Title: fmt.Sprintf(constants.DocumentTitle, "Contact"),
	//})
}
