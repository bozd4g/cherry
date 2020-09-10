package pageController

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/constants"
	"github.com/bozd4g/cherry/models"
	"github.com/bozd4g/cherry/services/mediumService"
	"github.com/gin-gonic/gin"
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
		Posts: controller.MediumService.GetPosts(),
	})
}

func (controller *pageController) postHandler(c *gin.Context) {
	posts := controller.MediumService.GetPosts()
	var selectedPost mediumService.PostDto
	for _, v := range posts {
		if v.Id == c.Param("id") {
			selectedPost = v
		}
	}

	c.HTML(http.StatusOK, "/post", models.PostDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, selectedPost.Title),
		Body:  template.HTML(selectedPost.Content),
	})
}

func (controller *pageController) aboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "/about", models.AboutDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "About"),
	})
}

func (controller *pageController) contactHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "/contact", models.ContactDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "Contact"),
	})
}
