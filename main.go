package main

import (
	"fmt"
	"github.com/bozd4g/cherry/routes"
	"github.com/bozd4g/cherry/utils"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	utils.LoadTemplates()
	r := routes.NewRouter()

	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(origins, headers, methods)(r))

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(fmt.Sprintf("Listening on localhost:%s", port))
	}
}
