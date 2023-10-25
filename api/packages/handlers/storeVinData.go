package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gc/db"
	"github.com/gc/helpers"
	"github.com/gc/types"
	"github.com/gc/vin"
)

// POST /storevindata
// store the specified vin data in the database
func PostStoreVINData(w http.ResponseWriter, r *http.Request) {
	// parse the query parameters from the complete request url
	v := r.URL.Query().Get("vin")
	if v == "" {
		// decode the request body
		var d types.VINData
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			helpers.ErrorResponse(&w, "An error occurred while decoding the request body, the body may not have been provided", err)
		} else {
			// store the provided vin data

			// connect to database
			dbconn, err := db.Connect()
			if err != nil {
				helpers.ErrorResponse(&w, "An error occurred while connecting to the database.", err)
			}

			// insert the data into the vindata table
			v, err := db.InsertVINData(dbconn, d)
			if err != nil {
				helpers.ErrorResponse(&w, "An error occurred while attempting to insert the data into the database", err)
			} else {
				// send back a success message
				helpers.MessageResponse(&w, fmt.Sprintf("Data for VIN %s was successfully inserted", v))
			}
		}
	} else {
		// look up the vin data
		d, err := vin.Lookup(v)
		if err != nil {
			helpers.ErrorResponse(&w, "An error occured while looking up the VIN data", err)
		} else {
			// make sure there's one and only one record for the specified vin
			if d.Count == 0 {
				// there are no records
				helpers.ErrorResponse(&w, "No data results were found for the specified VIN", nil)
			} else if d.Count > 1 {
				// there are two or more records
				helpers.ErrorResponse(&w, fmt.Sprintf("%d results were found for the specified VIN", d.Count), nil)
			} else {
				// there is a single record, store the retrieved data

				// connect to database
				dbconn, err := db.Connect()
				if err != nil {
					helpers.ErrorResponse(&w, "An error occurred while connecting to the database.", err)
				}

				// insert the data into the vindata table
				v, err := db.InsertVINData(dbconn, d.Results[0])
				if err != nil {
					helpers.ErrorResponse(&w, "An error occurred while attempting to insert the data into the database", err)
				} else {
					// send back a success message
					helpers.MessageResponse(&w, fmt.Sprintf("Data for VIN %s was successfully inserted", v))
				}
			}
		}
	}
}
