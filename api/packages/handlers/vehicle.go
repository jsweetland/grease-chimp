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
	var plateissuer string
	var platevalue string
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
	SELECT vin, plateissuer, platevalue, make, model, year, trim, package, nickname, colorname, colorhex FROM vehicles
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
			&plateissuer,
			&platevalue,
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
			ID:  id,
			VIN: vin,
			Plate: types.Plate{
				Issuer: plateissuer,
				Value:  platevalue,
			},
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

// POST /vehicle
// add or update data for a vehicle
func PostAddOrUpdateVehicle(w http.ResponseWriter, r *http.Request) {
	var v types.Vehicle
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		helpers.ErrorResponse(&w, "an error occurred decoding the request body", err)
	} else {
		// connect to database
		dbconn, err := db.Connect()
		if err != nil {
			helpers.ErrorResponse(&w, "An error occurred while connecting to the database.", err)
		}

		if v.ID == 0 {
			// the vehicle data does not have an id, insert the vehicle data
			id, err := db.InsertVehicleData(dbconn, v)
			if err != nil {
				helpers.ErrorResponse(&w, "An error occurred while inserting vehicle data.", err)
			}
			v.ID = id
		} else {
			// the vehicle data has an id, update the existing data
			err := db.UpdateVehicleData(dbconn, v)
			if err != nil {
				helpers.ErrorResponse(&w, "An error occurred while inserting vehicle data.", err)
			}
		}

		// send the data back in the response
		helpers.EnableCors(&w)
		helpers.SetJSONContentType(&w)
		json.NewEncoder(w).Encode(v)
	}
}
