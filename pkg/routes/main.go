package routes

import (
	"net/http"

	"github.com/kingbosman/url-shortener/pkg/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{encoded_url}", handlers.NewUrl)
	mux.HandleFunc("POST /new", handlers.NewUrl)

	return mux

}
