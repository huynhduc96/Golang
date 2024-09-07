package handlers

import (
	"database/Assignment/http-client/internal/constant"
	"encoding/json"
	"net/http"
)

// ErrorResponse struct to define the error response object
type ErrorResponse struct {
	Error string `json:"error"`
}

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithJSON sends a JSON response
func respondWithCookie(w http.ResponseWriter, name, value string) {
	cookie := &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/",
		MaxAge: constant.REDIS_TTL, // Cookie expiration in seconds
	}
	http.SetCookie(w, cookie)
}

// respondWithError sends an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}
