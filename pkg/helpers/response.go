package helpers

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, stringResponse string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]any{
		"statusCode": http.StatusOK,
		"message":    stringResponse,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)

	}
}

func ErrorResponse(w http.ResponseWriter, stringResponse string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]any{
		"statusCode": statusCode,
		"message":    stringResponse,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)

	}
}
