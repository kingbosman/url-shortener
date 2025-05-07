package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alextanhongpin/base62"
	"github.com/kingbosman/url-shortener/pkg/helpers"
)

type UrlRequest struct {
	URL string `json:"url"`
}

func NewUrl(w http.ResponseWriter, r *http.Request) {

	var newUrl UrlRequest
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), 500)
		return
	}

	createTableStmt := `
	CREATE TABLE IF NOT EXISTS urls (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	url TEXT
	)
	`
	db, err := helpers.Connect()
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), 500)
		return
	}
	defer db.Close()

	_, err = db.Exec(createTableStmt)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), 500)
		return
	}

	res, err := db.Exec("INSERT into urls (url) VALUES(?)", newUrl.URL)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), 500)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), 500)
		return
	}

	sampleId := base62.Encode(uint64(id))
	helpers.SuccessResponse(w, fmt.Sprintf("https://www.redirectdomain.com/%v", sampleId))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	decodedId := base62.Decode(strings.Split(r.RequestURI, "/")[1])

	db, err := helpers.Connect()
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), 500)
		return
	}

	var url string
	err = db.QueryRow("select url from urls where id = :id", decodedId).Scan(&url)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), 500)
		return
	}

	helpers.SuccessResponse(w, url)
}
