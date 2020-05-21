package routes

import (
	"fmt"
	"github.com/bozd4g/cherry/constants"
	"github.com/bozd4g/cherry/models"
	"github.com/bozd4g/cherry/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/p/{title}/{id}", postHandler)
	r.HandleFunc("/about", aboutHandler)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", models.IndexDocument{
		Title: fmt.Sprintf(constants.DocumentTitle, "Home"),
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
