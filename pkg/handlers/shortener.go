package handlers

import (
	"net/http"

	"github.com/kingbosman/url-shortener/pkg/helpers"
)

func NewUrl(w http.ResponseWriter, r *http.Request) {
	helpers.SuccessResponse(w, "todo posting new url")
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	// TODO: get variable from url
	helpers.SuccessResponse(w, "todo redirect to")
}
