package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gc/helpers"
	"github.com/gc/types"
)

// handles any unhandled routes
func GetBaseRoute(w http.ResponseWriter, r *http.Request) {
	m := types.SimpleMessage{
		Message: "no functionality is implemented at this endpoint",
	}

	helpers.EnableCors(&w)
	helpers.SetJSONContentType(&w)
	json.NewEncoder(w).Encode(m)
}
