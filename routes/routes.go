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
	"regexp"
	"strings"
	"time"
)

var memoryCache *cache.Cache

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/p/{id}/{title}", postHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)

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
	rss := getRss()
	var selectedRss models.ItemDto
	for _, v := range rss.Items {
		if v.Id == fmt.Sprintf("p/%s", vars["id"]) {
			selectedRss = v
		}
	}

	utils.ExecuteTemplate(w, "post.html", models.PostDocument{
		Title:       fmt.Sprintf(constants.DocumentTitle, selectedRss.Title),
		Description: "Lorem ipsum dolor",
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "about.html", models.AboutDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "About"),
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "contact.html", models.ContactDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "Contact"),
	})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.StatusNotFound(w)
}

func getRss() models.RssDto {
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

	rssDto := models.RssDto{}
	err := mapstructure.Decode(rss, &rssDto)
	if err != nil {
		log.Fatal(err)
	}

	var colNumber = 4
	var rgx = regexp.MustCompile(`(http[s]?:\/\/)?([^\/\s]+\/)(.*)`)

	for i := 0; i < len(rssDto.Items); i++ {
		itemDto := &rssDto.Items[i]
		if len(itemDto.Categories) == 0 {
			rssDto.Items = append(rssDto.Items[:i], rssDto.Items[i+1:]...)
			i--
			continue
		}

		guidMatches := rgx.FindAllStringSubmatch(itemDto.Guid, -1)
		linkMatches := rgx.FindAllStringSubmatch(itemDto.Link, -1)

		itemDto.Id = guidMatches[0][3]
		itemDto.Link = fmt.Sprintf("%s/%s", guidMatches[0][3], strings.Replace(linkMatches[0][3], "/", "-", 10))
		itemDto.ClassName = fmt.Sprintf("col-md-%d", colNumber)

		if i == 0 {
			colNumber = 8
		} else if i == 4 {
			colNumber = 8
		} else {
			colNumber = 4
		}
	}

	return rssDto
}
