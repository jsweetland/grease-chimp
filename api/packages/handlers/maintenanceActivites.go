package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gc/db"
	"github.com/gc/helpers"
	"github.com/gc/types"
)

// GET /maintenance/records/{vehicleID}
// get maintenance records for the vehicle with the specified vehicleID
func GetMaintenanceActivities(w http.ResponseWriter, r *http.Request) {
	var id int
	var activity string

	// connect to database
	db, err := db.Connect()
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while connecting to the database.", err)
	}

	// query the database for vehicles
	q := `SELECT id, activity FROM maintenanceactivities`
	rows, err := db.Query(q)
	if err != nil {
		helpers.ErrorResponse(&w, "An error occurred while querying the maintenanceactivities table in the database.", err)
	}
	defer rows.Close()

	// load the data
	var activities []types.MaintenanceActivity
	for rows.Next() {
		var v types.MaintenanceActivity
		rows.Scan(
			&id,
			&activity,
		)
		v = types.MaintenanceActivity{
			ID:       id,
			Activity: activity,
		}
		activities = append(activities, v)
	}

	// send the data back in the response
	helpers.EnableCors(&w)
	helpers.SetJSONContentType(&w)
	json.NewEncoder(w).Encode(activities)
}
