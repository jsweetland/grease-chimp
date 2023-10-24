package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"gc/db"
	"gc/vin"
)

// -----   Data Types   -----

// type for a simple message response
type SimpleMessage struct {
	Message string `json:"message"`
}

// type for an error message response
type ErrorMessage struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// type for the about info for the service
type About struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// type for the vehicle color data
type Color struct {
	Name string `json:"name,omitempty"`
	Hex  string `json:"hex,omitempty"`
}

// type for the vehicle data
type Vehicle struct {
	ID       int    `json:"id"`
	VIN      string `json:"vin,omitempty"`
	Make     string `json:"make,omitempty"`
	Model    string `json:"model,omitempty"`
	Year     int    `json:"year,omitempty"`
	Trim     string `json:"trim,omitempty"`
	Package  string `json:"package,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Color    Color  `json:"color,omitempty"`
}

// -----   Static Values   -----

var port = "10000"

// set application about info
var AboutInfo = About{
	Name:    "Grease Chimp API",
	Version: "0.0.0",
}

// -----   HTTP Helper Functions   -----

// open up cors access control
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// set json response content type header
func setJSONContentType(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

// log the name of the current handler
func logHandler(h string, m string) {
	fmt.Printf("handler - %s %s\n", m, h)
}

// return a message response
func messageResponse(w *http.ResponseWriter, m string) {
	r := SimpleMessage{
		Message: m,
	}

	enableCors(w)
	setJSONContentType(w)
	json.NewEncoder(*w).Encode(r)
}

// return an error response
func errorResponse(w *http.ResponseWriter, m string, err error) {
	r := ErrorMessage{
		Message: m,
		Error:   err,
	}

	enableCors(w)
	setJSONContentType(w)
	(*w).WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(*w).Encode(r)
}

// send method not allowed response
func methodNotAllowedResponse(w *http.ResponseWriter, h string) {
	s := "Method not allowed"
	fmt.Printf("  [%s] %s\n", h, s)
	http.Error((*w), s, http.StatusMethodNotAllowed)
}

// -----   Route Handlers   -----

// handles any unhandled routes
func getBaseRoute(w http.ResponseWriter, r *http.Request) {
	h := "getBaseRoute"
	logHandler(h, r.Method)

	s := "no functionality is implemented at this endpoint"
	m := SimpleMessage{
		Message: s,
	}

	fmt.Printf("  [%s] %s\n", h, s)

	enableCors(&w)
	setJSONContentType(&w)
	json.NewEncoder(w).Encode(m)
}

// GET /about
// returns the about info for the application
func getAbout(w http.ResponseWriter, r *http.Request) {
	h := "getAbout"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodGet:
		enableCors(&w)
		setJSONContentType(&w)
		json.NewEncoder(w).Encode(AboutInfo)
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// GET /vehicles
// returns the vehicles stored in the database
func getVehicles(w http.ResponseWriter, r *http.Request) {
	h := "getVehicles"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodGet:
		var id int
		var vin string
		var make string
		var model string
		var year int
		var trim string
		var trimpackage string
		var nickname string
		var colorname string
		var colorhex string
		var vehicles = []Vehicle{}

		// connect to database
		db, err := db.Connect()
		if err != nil {
			errorResponse(&w, "An error occurred while connecting to the database.", err)
		}

		// query the database for vehicles
		q := "SELECT id, vin, make, model, year, trim, package, nickname, colorname, colorhex FROM vehicles"
		rows, err := db.Query(q)
		defer rows.Close()
		if err != nil {
			errorResponse(&w, "An error occurred while querying the vehicles table in the database.", err)
		}

		// load the rows of data into an array
		for rows.Next() {
			rows.Scan(
				&id,
				&vin,
				&make,
				&model,
				&year,
				&trim,
				&trimpackage,
				&nickname,
				&colorname,
				&colorhex,
			)
			v := Vehicle{
				ID:       id,
				VIN:      vin,
				Make:     make,
				Model:    model,
				Year:     year,
				Trim:     trim,
				Package:  trimpackage,
				Nickname: nickname,
				Color: Color{
					Name: colorname,
					Hex:  colorhex,
				},
			}
			vehicles = append(vehicles, v)
		}

		// send the data back in the response
		enableCors(&w)
		setJSONContentType(&w)
		json.NewEncoder(w).Encode(vehicles)
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// GET /vinlookup?vin={vin}
// look up the data for the specified vin
func getVINLookup(w http.ResponseWriter, r *http.Request) {
	h := "getVinLookup"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodGet:
		// parse the complete url from the request
		u, err := url.Parse(r.URL.String())
		if err != nil {
			errorResponse(&w, "An error occurred while parsing the URL for VIN lookup", err)
		}

		// parse the query parameters from the complete request url
		q, _ := url.ParseQuery(u.RawQuery)

		// extract the vin from the query parameters
		v := q["vin"][0]
		fmt.Printf("  [%s] vin = %s\n", h, v)

		// look up the vin data
		d, err := vin.Lookup(v)
		if err != nil {
			errorResponse(&w, "An error occured while looking up the VIN data", err)
		}

		// return the vin data
		enableCors(&w)
		setJSONContentType(&w)
		json.NewEncoder(w).Encode(d)
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// POST /storevindata
// store the specified vin data in the database
func postStoreVINData(w http.ResponseWriter, r *http.Request) {
	h := "postStoreVINData"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodPost:
		// parse the complete url from the request
		u, err := url.Parse(r.URL.String())
		if err != nil {
			errorResponse(&w, "An error occurred while parsing the URL for VIN lookup", err)
		}

		// parse the query parameters from the complete request url
		q, _ := url.ParseQuery(u.RawQuery)
		if len(q) == 0 || q["vin"] == nil {
			fmt.Printf("  [%s] no vin provided, will use vin data in request body\n", h)

			// decode the request body
			var d vin.LookupResult
			err = json.NewDecoder(r.Body).Decode(&d)
			if err != nil {
				errorResponse(&w, "An error occurred while decoding the request body, the body may not have been provided", err)
			} else {
				// store the provided vin data

				// connect to database
				dbconn, err := db.Connect()
				if err != nil {
					errorResponse(&w, "An error occurred while connecting to the database.", err)
				}

				// insert the data into the vindata table
				v, err := db.InsertVINData(dbconn, d)
				if err != nil {
					errorResponse(&w, "An error occurred while attempting to insert the data into the database", err)
				} else {
					// send back a success message
					fmt.Printf("  [%s] data for VIN %s successfully inserted", h, v)
					messageResponse(&w, fmt.Sprintf("Data for VIN %s was successfully inserted", v))
				}
			}
		} else {
			// extract the vin from the query parameters
			v := q["vin"][0]
			fmt.Printf("  [%s] vin %s provided, will look up data to store\n", h, v)

			// look up the vin data
			d, err := vin.Lookup(v)
			if err != nil {
				errorResponse(&w, "An error occured while looking up the VIN data", err)
			} else {
				// make sure there's one and only one record for the specified vin
				if d.Count == 0 {
					// there are no records
					errorResponse(&w, "No data results were found for the specified VIN", nil)
				} else if d.Count > 1 {
					// there are two or more records
					errorResponse(&w, fmt.Sprintf("%d results were found for the specified VIN", d.Count), nil)
				} else {
					// there is a single record, store the retrieved data
					fmt.Printf("  [%s] data for VIN %s successfully retrieved", h, v)

					// connect to database
					dbconn, err := db.Connect()
					if err != nil {
						errorResponse(&w, "An error occurred while connecting to the database.", err)
					}

					// insert the data into the vindata table
					v, err := db.InsertVINData(dbconn, d.Results[0])
					if err != nil {
						errorResponse(&w, "An error occurred while attempting to insert the data into the database", err)
					} else {
						// send back a success message
						fmt.Printf("  [%s] data for VIN %s successfully inserted", h, v)
						messageResponse(&w, fmt.Sprintf("Data for VIN %s was successfully inserted", v))
					}
				}
			}
		}
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// handle the request routes
func handleRequests() {
	http.HandleFunc("/", getBaseRoute)
	http.HandleFunc("/about", getAbout)
	http.HandleFunc("/vehicles", getVehicles)
	http.HandleFunc("/vinlookup", getVINLookup)
	http.HandleFunc("/storevindata", postStoreVINData)

	fmt.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// main
func main() {
	handleRequests()
}
