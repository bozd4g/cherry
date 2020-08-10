package main

import (
	"fmt"
	"github.com/bozd4g/cherry/controllers"
	"os"
)

func main() {
	r := controllers.New()

	r.InitRoutes()
	r.InitMiddlewares()
	r.InitTemplates()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Get().Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
