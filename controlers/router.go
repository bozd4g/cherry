package controlers

import (
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/controlers/mediumController"
	"github.com/bozd4g/cherry/controlers/pageController"
	"github.com/bozd4g/cherry/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type router struct {
	Router *mux.Router
}

type IRouter interface {
	InitRoutes() http.Handler
	InitMiddlewares() http.Handler
}

func New() IRouter {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return &router{Router: r}
}

func (router *router) InitRoutes() http.Handler {
	router.Router.NotFoundHandler = http.HandlerFunc(router.notFoundHandler)

	memoryCache := caching.New()
	router.Router = pageController.New(router.Router, memoryCache).Init()
	router.Router = mediumController.New(router.Router, memoryCache).Init()

	return router.Router
}

func (router *router) InitMiddlewares() http.Handler {
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(origins, headers, methods)(router.Router)
}

func (router *router) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.StatusNotFound(w)
}
