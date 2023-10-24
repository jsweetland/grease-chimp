package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/gc/types"
)

// open up cors access control
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// set json response content type header
func SetJSONContentType(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

// return a message response
func MessageResponse(w *http.ResponseWriter, m string) {
	r := types.SimpleMessage{
		Message: m,
	}

	EnableCors(w)
	SetJSONContentType(w)
	json.NewEncoder(*w).Encode(r)
}

// return an error response
func ErrorResponse(w *http.ResponseWriter, m string, err error) {
	r := types.ErrorMessage{
		Message: m,
		Error:   err,
	}

	EnableCors(w)
	SetJSONContentType(w)
	(*w).WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(*w).Encode(r)
}
