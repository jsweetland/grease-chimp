package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gc/db"
	"github.com/go-chi/chi/v5"

	"github.com/gc/helpers"
	"github.com/gc/types"
)

// GET /vehicle/{id}
// get data for the vehicle with the specified id
func GetVehicleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while parsing the vehicle ID from the request parameters.", err)
	}

	var vin string
	var make string
	var model string
	var year int
	var trim string
	var trimpackage string
	var nickname string
	var colorname string
	var colorhex string

	// connect to database
	db, err := db.Connect()
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while connecting to the database.", err)
	}

	// query the database for vehicles
	q := `
	SELECT vin, make, model, year, trim, package, nickname, colorname, colorhex FROM vehicles
	WHERE id = $1`
	rows, err := db.Query(q, id)
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while querying the vehicles table in the database.", err)
	}
	defer rows.Close()

	// load the data
	var v types.Vehicle
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
		v = types.Vehicle{
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
	}

	// send the data back in the response
	helpers.EnableCors(&w)
	helpers.SetJSONContentType(&w)
	json.NewEncoder(w).Encode(v)
}