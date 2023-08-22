package main

import (
    "fmt"
    "log"
    "net/http"
		"encoding/json"
		"database/sql"
		_ "github.com/lib/pq"
)

var port = "10000"
var dbConnStr = "postgresql://gcuser:gcpass@localhost/gc?sslmode=disable"

type SimpleMessage struct {
	Message string `json:"message"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Error error `json:"error"`
}

type About struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

var AboutInfo = About{
	Name: "Grease Chimp API",
	Version: "0.0.0",
}

type Color struct {
	Name string `json:"name,omitempty"`
	Hex string `json:"hex,omitempty"`
}

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getBaseRoute(w http.ResponseWriter, r *http.Request){
	fmt.Println("in handler: getBaseRoute")
	
	m := SimpleMessage{
		Message: "no functionality is implemented at this endpoint",
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func getAbout(w http.ResponseWriter, r *http.Request){
	fmt.Println("in handler: getAbout")

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AboutInfo)
}

func getVehicles(w http.ResponseWriter, r *http.Request){
	fmt.Println("in handler: getVehicles")

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
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		m := ErrorMessage{
			Message: "An error occurred while connecting to the database.",
			Error: err,
		}

		enableCors(&w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	}

	q := "SELECT vin, make, model, year, trim, package, nickname, colorname, colorhex FROM vehicles"
	rows, err := db.Query(q)
	defer rows.Close()
	if err != nil {
		m := ErrorMessage{
			Message: "An error occurred while querying the vehicles table in the database.",
			Error: err,
		}

		enableCors(&w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	}
	
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
	
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}

func handleRequests() {
	http.HandleFunc("/", getBaseRoute)
	http.HandleFunc("/about", getAbout)
	http.HandleFunc("/vehicles", getVehicles)

	fmt.Println(fmt.Sprintf("listening on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	handleRequests()
}
