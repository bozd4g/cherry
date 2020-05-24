package routes

import (
	"fmt"
	"github.com/bozd4g/cherry/constants"
	"github.com/bozd4g/cherry/models"
	"github.com/bozd4g/cherry/utils"
	"github.com/bozd4g/go-http-client/client"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"os"
	"time"
)

var memoryCache *cache.Cache

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/p/{title}/{id}", postHandler)
	r.HandleFunc("/about", aboutHandler)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	memoryCache = cache.New(8*time.Hour, 10*time.Hour)
	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", models.IndexDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "Home"),
		Rss:   getRss(),
	})
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	utils.ExecuteTemplate(w, "post.html", models.PostDocument{
		Title:       fmt.Sprintf(constants.DocumentTitle, vars["title"]),
		Description: "Lorem ipsum dolor",
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "about.html", models.AboutDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "About"),
	})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.StatusNotFound(w)
}

func getRss() models.Rss {
	apiBaseUrl := os.Getenv("API_BASE_URL")
	apiGetMethod := os.Getenv("API_GET_METHOD")

	rss := models.Rss{}
	if apiBaseUrl != "" && apiGetMethod != "" {
		var err error
		if rssCache, isRssCacheExist := memoryCache.Get("rssData"); !isRssCacheExist {
			httpClient := client.HttpClient{BaseUrl: apiBaseUrl}
			response := httpClient.Get(apiGetMethod)

			if response.IsSuccess {
				err = mapstructure.Decode(response.Data, &rss)
			} else {
				log.Println("Error: " + response.Message)
			}

			memoryCache.Set("rssData", rss, cache.DefaultExpiration)
		} else {
			err = mapstructure.Decode(rssCache, &rss)
			log.Println("Data came from cache")
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return rss
}
