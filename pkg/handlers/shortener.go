package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alextanhongpin/base62"
	"github.com/kingbosman/url-shortener/pkg/helpers"
)

func NewUrl(w http.ResponseWriter, r *http.Request) {
	sampleId := base62.Encode(1232)
	helpers.SuccessResponse(w, fmt.Sprintf("encoded id: %v", sampleId))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	decoded := base62.Decode(strings.Split(r.RequestURI, "/")[1])
	helpers.SuccessResponse(w, fmt.Sprintf("Redirecting to url with id %v", decoded))
}
