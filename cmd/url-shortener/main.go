package main

import (
	"log"
	"net/http"

	"github.com/kingbosman/url-shortener/pkg/routes"
)

func main() {
	routes := routes.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", routes))
}
