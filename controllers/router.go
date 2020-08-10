package controllers

import (
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/controllers/mediumController"
	"github.com/bozd4g/cherry/controllers/pageController"
	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
	"net/http"
	"path/filepath"
)

type Router struct {
	Engine *gin.Engine
}

type IRouter interface {
	Get() *gin.Engine
	InitRoutes()
	InitMiddlewares()
	InitTemplates()
}

func New() IRouter {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*/*.html")
	engine.Static("static", "./static")

	engine.NoRoute(func(context *gin.Context) {
		context.Data(http.StatusNotFound, "text/plain", []byte("404 - Not found!"))
	})

	return &Router{Engine: engine}
}

func (router *Router) Get() *gin.Engine {
	return router.Engine
}

func (router *Router) InitRoutes() {
	memoryCache := caching.New()

	pageController.New(router.Engine, memoryCache).Init()
	mediumController.New(router.Engine, memoryCache).Init()
}

func (router *Router) InitMiddlewares() {
	router.Engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

func (router *Router) InitTemplates() {
	render := eztemplate.New()
	render.TemplatesDir = filepath.Join("templates")
	render.Layout = "/layout/_base"
	router.Engine.HTMLRender = render.Init()
}
