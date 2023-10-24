package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gc/db"
	"github.com/gc/helpers"
	"github.com/gc/types"
)

// GET /vehicles
// returns the vehicles stored in the database
func GetVehicles(w http.ResponseWriter, r *http.Request) {
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
	var vehicles = []types.Vehicle{}

	// connect to database
	db, err := db.Connect()
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while connecting to the database.", err)
	}

	// query the database for vehicles
	q := "SELECT id, vin, make, model, year, trim, package, nickname, colorname, colorhex FROM vehicles"
	rows, err := db.Query(q)
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while querying the vehicles table in the database.", err)
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
		v := types.Vehicle{
			ID:       id,
			VIN:      vin,
			Make:     make,
			Model:    model,
			Year:     year,
			Trim:     trim,
			Package:  trimpackage,
			Nickname: nickname,
			Color: types.Color{
				Name: colorname,
				Hex:  colorhex,
			},
		}
		vehicles = append(vehicles, v)
	}

	// send the data back in the response
	helpers.EnableCors(&w)
	helpers.SetJSONContentType(&w)
	json.NewEncoder(w).Encode(vehicles)
}
