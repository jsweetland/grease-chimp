package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gc/helpers"
	"github.com/gc/types"
)

var AboutInfo = types.About{
	Name:    "Grease Chimp API",
	Version: "0.0.0",
}

// GET /about
// returns the about info for the application
func GetAbout(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	helpers.SetJSONContentType(&w)
	json.NewEncoder(w).Encode(AboutInfo)
}
