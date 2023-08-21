package main

import (
    "fmt"
    "log"
    "net/http"
		"encoding/json"
)

var port = "10000"

type SimpleMessage struct {
	Message string `json:"message"`
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

var Vehicles []Vehicle

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
	
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(Vehicles)
}

func handleRequests() {
	http.HandleFunc("/", getBaseRoute)
	http.HandleFunc("/about", getAbout)
	http.HandleFunc("/vehicles", getVehicles)

	fmt.Println(fmt.Sprintf("listening on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	Vehicles = []Vehicle{
		Vehicle{
			VIN: "abc",
			Make: "Jeep",
			Model: "Wrangler Unlimited",
			Year: 2020,
			Trim: "Sport",
			Package: "Willys",
			Nickname: "Junebug",
			Color: Color{
				Name: "Hellayella",
				Hex: "fdb93c",
			},
		},
		Vehicle{
			VIN: "def",
			Make: "Toyota",
			Model: "Sienna",
			Year: 2013,
			Trim: "Limited",
		},
	}
		
	handleRequests()
}
