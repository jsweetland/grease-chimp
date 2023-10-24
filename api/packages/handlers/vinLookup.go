package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gc/helpers"
	"github.com/gc/vin"
)

// GET /vinlookup?vin={vin}
// look up the data for the specified vin
func GetVINLookup(w http.ResponseWriter, r *http.Request) {
	// get the vin value from the parameter
	v := r.URL.Query().Get("vin")
	if v == "" {
		helpers.ErrorResponse(&w, "No VIN value was provided", nil)
	} else {
		// look up the vin data
		d, err := vin.Lookup(v)
		if err != nil {
			helpers.ErrorResponse(&w, "An error occured while looking up the VIN data", err)
		}

		// return the vin data
		helpers.EnableCors(&w)
		helpers.SetJSONContentType(&w)
		json.NewEncoder(w).Encode(d)
	}
}
