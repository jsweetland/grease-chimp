package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gc/db"
	"github.com/gc/helpers"
	"github.com/gc/types"
	"github.com/go-chi/chi/v5"
)

// GET /maintenance/records/{vehicleID}
// get maintenance records for the vehicle with the specified vehicleID
func GetMaintenanceRecordsByVehicleID(w http.ResponseWriter, r *http.Request) {
	vehicleID, err := strconv.Atoi(chi.URLParam(r, "vehicleID"))
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while parsing the vehicle ID from the request parameters.", err)
	}

	var id int
	var activity string
	var datePerformed string
	var mileage int

	// connect to database
	db, err := db.Connect()
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while connecting to the database.", err)
	}

	// query the database for vehicles
	q := `
	SELECT maintenancerecords.id, maintenanceactivities.activity, maintenancerecords.dateperformed, maintenancerecords.mileage
	FROM maintenancerecords
	INNER JOIN maintenanceactivities ON maintenancerecords.activityid=maintenanceactivities.id
	WHERE vehicleid = $1`
	rows, err := db.Query(q, vehicleID)
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while querying the maintenancerecords table in the database.", err)
	}
	defer rows.Close()

	// load the data
	var records []types.MaintenanceRecord
	for rows.Next() {
		var v types.MaintenanceRecord
		rows.Scan(
			&id,
			&activity,
			&datePerformed,
			&mileage,
		)
		v = types.MaintenanceRecord{
			ID:            id,
			VehicleID:     vehicleID,
			Activity:      activity,
			DatePerformed: datePerformed,
			Mileage:       mileage,
		}
		records = append(records, v)
	}

	// send the data back in the response
	helpers.EnableCors(&w)
	helpers.SetJSONContentType(&w)
	json.NewEncoder(w).Encode(records)
}
