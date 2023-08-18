package main

import (
    "fmt"
    "log"
    "net/http"
		"encoding/json"
)

var port = "10000"

type About struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

var AboutInfo = About{
	Name: "Grease Chimp API",
	Version: "0.0.0",
}

type Vehicle struct {
	Make string `json:"make"`
	Model string `json:"model"`
	Year int `json:"year"`
	Nickname string `json:"nickname"`
}

var Vehicles []Vehicle

func getBaseRoute(w http.ResponseWriter, r *http.Request){
	fmt.Println("in handler: getBaseRoute")
	json.NewEncoder(w).Encode({
		message: "no functionality is implemented at this endpoint",
	})
}

func getAbout(w http.ResponseWriter, r *http.Request){
	fmt.Println("in handler: getAbout")
	json.NewEncoder(w).Encode(AboutInfo)
}

func getVehicles(w http.ResponseWriter, r *http.Request){
	fmt.Println("in handler: getVehicles")
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
			Make: "Jeep",
			Model: "Wrangler Unlimited",
			Year: 2020,
			Nickname: "Junebug",
		},
	}
		
	handleRequests()
}
