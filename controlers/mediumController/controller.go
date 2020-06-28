package mediumController

import (
	"encoding/json"
	"github.com/bozd4g/cherry/caching"
	"github.com/bozd4g/cherry/proxy/mediumProxy/mediumProxyDtos"
	"github.com/bozd4g/cherry/services/mediumService"
	"github.com/gorilla/mux"
	"net/http"
)

type mediumController struct {
	Routes        *mux.Router
	MediumService mediumService.IMediumService
}

type IMediumController interface {
	Init() *mux.Router
}

func New(routes *mux.Router, memoryCache caching.IMemoryCache) IMediumController {
	mediumService := mediumService.New(memoryCache)
	return &mediumController{Routes: routes, MediumService: mediumService}
}

func (mc *mediumController) Init() *mux.Router {
	mc.Routes.HandleFunc("/api/medium/feed", mc.feedHandler).Methods(http.MethodGet)
	return mc.Routes
}

func (mc *mediumController) feedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	rss := mc.MediumService.GetRss()
	response, err := json.Marshal(rss)
	if err != nil {
		emptyResponse, _ := json.Marshal(mediumProxyDtos.RssDto{
			Status: "",
			Feed:   mediumProxyDtos.FeedDto{},
			Items:  []mediumProxyDtos.ItemDto{},
		})
		response = emptyResponse
	}

	_, err = w.Write(response)
	if err != nil {
		panic(err)
	}
}
