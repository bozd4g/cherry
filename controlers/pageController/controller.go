package pageController

import (
	"fmt"
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/constants"
	"github.com/bozd4g/cherry/models"
	"github.com/bozd4g/cherry/proxy/mediumProxy/mediumProxyDtos"
	"github.com/bozd4g/cherry/services/mediumService"
	"github.com/bozd4g/cherry/utils"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type pageController struct {
	Routes        *mux.Router
	MediumService mediumService.IMediumService
}

type IPageController interface {
	Init() *mux.Router
}

func New(routes *mux.Router, memoryCache caching.IMemoryCache) IPageController {
	mediumService := mediumService.New(memoryCache)
	return &pageController{Routes: routes, MediumService: mediumService}
}

func (pc *pageController) Init() *mux.Router {
	pc.Routes.HandleFunc("/", pc.indexHandler).Methods(http.MethodGet)
	pc.Routes.HandleFunc("/p/{id}/{title}", pc.postHandler).Methods(http.MethodGet)
	pc.Routes.HandleFunc("/about", pc.aboutHandler).Methods(http.MethodGet)
	pc.Routes.HandleFunc("/contact", pc.contactHandler).Methods(http.MethodGet)

	return pc.Routes
}

func (pc *pageController) indexHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", models.IndexDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "Home"),
		Rss:   pc.MediumService.GetRss(),
	})
}

func (pc *pageController) postHandler(w http.ResponseWriter, r *http.Request) {
	rss := pc.MediumService.GetRss()
	var selectedRss mediumProxyDtos.ItemDto
	for _, v := range rss.Items {
		if v.Id == fmt.Sprintf("p/%s", mux.Vars(r)["id"]) {
			selectedRss = v
		}
	}

	utils.ExecuteTemplate(w, "post.html", models.PostDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, selectedRss.Title),
		Body:  template.HTML(selectedRss.Content),
	})
}

func (pc *pageController) aboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "about.html", models.AboutDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "About"),
	})
}

func (pc *pageController) contactHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "contact.html", models.ContactDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "Contact"),
	})
}
