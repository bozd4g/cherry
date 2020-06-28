package main

import (
	"fmt"
	"github.com/bozd4g/cherry/controlers"
	"github.com/bozd4g/cherry/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	utils.LoadTemplates()
	r := controlers.New()

	r.InitRoutes()
	middlewares := r.InitMiddlewares()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), middlewares)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(fmt.Sprintf("Listening on localhost:%s", port))
	}
}
