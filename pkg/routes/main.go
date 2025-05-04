package routes

import (
	"net/http"

	"github.com/kingbosman/url-shortener/pkg/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{encoded_url}", handlers.Redirect)
	mux.HandleFunc("POST /", handlers.NewUrl)

	return mux

}
