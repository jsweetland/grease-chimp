package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gc/db"
	"github.com/gc/vin"
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
	m := SimpleMessage{
		Message: "no functionality is implemented at this endpoint",
	}

	enableCors(&w)
	setJSONContentType(&w)
	json.NewEncoder(w).Encode(m)
}

// GET /about
// returns the about info for the application
func getAbout(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	setJSONContentType(&w)
	json.NewEncoder(w).Encode(AboutInfo)
}

// GET /vehicles
// returns the vehicles stored in the database
func getVehicles(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		errorResponse(&w, "An error occurred while querying the vehicles table in the database.", err)
	}
	defer rows.Close()

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
}

// GET /vinlookup?vin={vin}
// look up the data for the specified vin
func getVINLookup(w http.ResponseWriter, r *http.Request) {
	// get the vin value from the parameter
	v := r.URL.Query().Get("vin")
	if v == "" {
		errorResponse(&w, "No VIN value was provided", nil)
	} else {
		// look up the vin data
		d, err := vin.Lookup(v)
		if err != nil {
			errorResponse(&w, "An error occured while looking up the VIN data", err)
		}

		// return the vin data
		enableCors(&w)
		setJSONContentType(&w)
		json.NewEncoder(w).Encode(d)
	}
}

// POST /storevindata
// store the specified vin data in the database
func postStoreVINData(w http.ResponseWriter, r *http.Request) {
	// parse the query parameters from the complete request url
	v := r.URL.Query().Get("vin")
	if v == "" {
		// decode the request body
		var d vin.LookupResult
		err := json.NewDecoder(r.Body).Decode(&d)
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
				messageResponse(&w, fmt.Sprintf("Data for VIN %s was successfully inserted", v))
			}
		}
	} else {
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
					messageResponse(&w, fmt.Sprintf("Data for VIN %s was successfully inserted", v))
				}
			}
		}
	}
}

func startServer() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", getBaseRoute)
	r.Get("/about", getAbout)
	r.Get("/vehicles", getVehicles)
	r.Get("/vinlookup", getVINLookup)
	r.Post("/storevindata", postStoreVINData)

	p := 10000
	fmt.Printf("Listening on port %d\n", p)
	http.ListenAndServe(fmt.Sprintf(":%d", p), r)
}

func main() {
	startServer()
}
