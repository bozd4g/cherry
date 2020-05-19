package main

import (
	"fmt"
	"github.com/bozd4g/cherry/routes"
	"github.com/bozd4g/cherry/utils"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()

	port := 8080
	http.Handle("/", r)

	fmt.Println(fmt.Sprintf("Listening on localhost:%d", port))

	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(origins, headers, methods)(r)))
}
