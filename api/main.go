package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"encoding/json"

	"gc/vin"
	"gc/db"
)


// -----   Data Types   -----

// type for a simple message response
type SimpleMessage struct {
	Message string `json:"message"`
}

// type for an error message response
type ErrorMessage struct {
	Message string `json:"message"`
	Error error `json:"error"`
}

// type for the about info for the service
type About struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

//  type for the vehicle color data
type Color struct {
	Name string `json:"name,omitempty"`
	Hex string `json:"hex,omitempty"`
}

// type for the vehicle data
type Vehicle struct {
	VIN string `json:"vin,omitempty"`
	Make string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
	Year int `json:"year,omitempty"`
	Trim string `json:"trim,omitempty"`
	Package string `json:"package,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Color Color `json:"color,omitempty"`
}


// -----   Static Values   -----

var port = "10000"

// set application about info
var AboutInfo = About{
	Name: "Grease Chimp API",
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
func logHandler(h string) {
	fmt.Printf("handler - %s\n", h)
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
		Error: err,
	}

	enableCors(w)
	setJSONContentType(w)
	json.NewEncoder(*w).Encode(r)
}


// -----   Route Handlers   -----

// handles any unhandled routes
func getBaseRoute(w http.ResponseWriter, r *http.Request){
	h := "getBaseRoute"
	logHandler(h)
	
	m := SimpleMessage{
		Message: "no functionality is implemented at this endpoint",
	}

	enableCors(&w)
	setJSONContentType(&w)
	json.NewEncoder(w).Encode(m)
}

// GET /about
// returns the about info for the application
func getAbout(w http.ResponseWriter, r *http.Request){
	h := "getAbout"
	logHandler(h)

	enableCors(&w)
	setJSONContentType(&w)
	json.NewEncoder(w).Encode(AboutInfo)
}

// GET /vehicles
// returns the vehicles stored in the database
func getVehicles(w http.ResponseWriter, r *http.Request){
	h := "getVehicles"
	logHandler(h)

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
	q := "SELECT vin, make, model, year, trim, package, nickname, colorname, colorhex FROM vehicles"
	rows, err := db.Query(q)
	defer rows.Close()
	if err != nil {
		errorResponse(&w, "An error occurred while querying the vehicles table in the database.", err)
	}
	
	// load the rows of data into an array
	for rows.Next() {
		rows.Scan(
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
			VIN: vin,
			Make: make,
			Model: model,
			Year: year,
			Trim: trim,
			Package: trimpackage,
			Nickname: nickname,
			Color: Color{
				Name: colorname,
				Hex: colorhex,
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
	h := "getVinLookup"
	logHandler(h)

	// parse the complete url from the request
	u, err := url.Parse(r.URL.String())
	if err != nil {
		errorResponse(&w, "An error occurred while parsing the URL for VIN lookup", err)
	}

	// parse the query parameters from the complete request url
	q, _ := url.ParseQuery(u.RawQuery)

	// extract the vin from the query parameters
	v := q["vin"][0]
	fmt.Printf("   [%s] vin = %s\n", h, v)

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

// handle the request routes
func handleRequests() {
	http.HandleFunc("/", getBaseRoute)
	http.HandleFunc("/about", getAbout)
	http.HandleFunc("/vehicles", getVehicles)
	http.HandleFunc("/vinlookup", getVINLookup)

	fmt.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// main
func main() {
	handleRequests()
}
