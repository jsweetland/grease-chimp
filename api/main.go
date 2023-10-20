package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"gc/db"
	"gc/vin"
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
	h := "getBaseRoute"
	logHandler(h, r.Method)

	s := "no functionality is implemented at this endpoint"
	m := SimpleMessage{
		Message: s,
	}

	fmt.Printf("  [%s] %s\n", h, s)

	enableCors(&w)
	setJSONContentType(&w)
	json.NewEncoder(w).Encode(m)
}

// GET /about
// returns the about info for the application
func getAbout(w http.ResponseWriter, r *http.Request) {
	h := "getAbout"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodGet:
		enableCors(&w)
		setJSONContentType(&w)
		json.NewEncoder(w).Encode(AboutInfo)
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// GET /vehicles
// returns the vehicles stored in the database
func getVehicles(w http.ResponseWriter, r *http.Request) {
	h := "getVehicles"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodGet:
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
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// GET /vinlookup?vin={vin}
// look up the data for the specified vin
func getVINLookup(w http.ResponseWriter, r *http.Request) {
	h := "getVinLookup"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodGet:
		// parse the complete url from the request
		u, err := url.Parse(r.URL.String())
		if err != nil {
			errorResponse(&w, "An error occurred while parsing the URL for VIN lookup", err)
		}

		// parse the query parameters from the complete request url
		q, _ := url.ParseQuery(u.RawQuery)

		// extract the vin from the query parameters
		v := q["vin"][0]
		fmt.Printf("  [%s] vin = %s\n", h, v)

		// look up the vin data
		d, err := vin.Lookup(v)
		if err != nil {
			errorResponse(&w, "An error occured while looking up the VIN data", err)
		}

		// return the vin data
		enableCors(&w)
		setJSONContentType(&w)
		json.NewEncoder(w).Encode(d)
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// POST /storevindata
// store the specified vin data in the database
func postStoreVINData(w http.ResponseWriter, r *http.Request) {
	h := "postStoreVINData"
	logHandler(h, r.Method)

	switch r.Method {
	case http.MethodPost:
		// parse the complete url from the request
		u, err := url.Parse(r.URL.String())
		if err != nil {
			errorResponse(&w, "An error occurred while parsing the URL for VIN lookup", err)
		}

		// parse the query parameters from the complete request url
		q, _ := url.ParseQuery(u.RawQuery)
		if len(q) == 0 || q["vin"] == nil {
			fmt.Printf("  [%s] no vin provided, will use vin data in request body\n", h)

			var d vin.LookupResult
			err = json.NewDecoder(r.Body).Decode(&d)
			if err != nil {
				errorResponse(&w, "An error occurred while decoding the request body, the body may not have been provided", err)
			} else {
				// store the provided vin data

				// connect to database
				db, err := db.Connect()
				if err != nil {
					errorResponse(&w, "An error occurred while connecting to the database.", err)
				}

				s := `
				INSERT INTO vindata (
					abs,
					activesafetysysnote,
					adaptivecruisecontrol,
					adaptivedrivingbeam,
					adaptiveheadlights,
					additionalerrortext,
					airbagloccurtain,
					airbaglocfront,
					airbaglocknee,
					airbaglocseatcushion,
					airbaglocside,
					autoreversesystem,
					automaticpedestrianalertingsound,
					axleconfiguration,
					axles,
					baseprice,
					batterya,
					batterya_to,
					batterycells,
					batteryinfo,
					batterykwh,
					batterykwh_to,
					batterymodules,
					batterypacks,
					batterytype,
					batteryv,
					batteryv_to,
					bedlengthin,
					bedtype,
					blindspotintervention,
					blindspotmon,
					bodycabtype,
					bodyclass,
					brakesystemdesc,
					brakesystemtype,
					busfloorconfigtype,
					buslength,
					bustype,
					can_aacn,
					cib,
					cashforclunkers,
					chargerlevel,
					chargerpowerkw,
					coolingtype,
					curbweightlb,
					custommotorcycletype,
					daytimerunninglight,
					destinationmarket,
					displacementcc,
					displacementci,
					displacementl,
					doors,
					drivetype,
					driverassist,
					dynamicbrakesupport,
					edr,
					esc,
					evdriveunit,
					electrificationlevel,
					engineconfiguration,
					enginecycles,
					enginecylinders,
					enginehp,
					enginehp_to,
					enginekw,
					enginemanufacturer,
					enginemodel,
					entertainmentsystem,
					errorcode,
					errortext,
					forwardcollisionwarning,
					fuelinjectiontype,
					fueltypeprimary,
					fueltypesecondary,
					gcwr,
					gcwr_to,
					gvwr,
					gvwr_to,
					keylessignition,
					lanecenteringassistance,
					lanedeparturewarning,
					lanekeepsystem,
					lowerbeamheadlamplightsource,
					make,
					makeid,
					manufacturer,
					manufacturerid,
					model,
					modelid,
					modelyear,
					motorcyclechassistype,
					motorcyclesuspensiontype,
					ncsabodytype,
					ncsamake,
					ncsamapexcapprovedby,
					ncsamapexcapprovedon,
					ncsamappingexception,
					ncsamodel,
					ncsanote,
					nonlanduse,
					note,
					otherbusinfo,
					otherengineinfo,
					othermotorcycleinfo,
					otherrestraintsysteminfo,
					othertrailerinfo,
					parkassist,
					pedestrianautomaticemergencybraking,
					plantcity,
					plantcompanyname,
					plantcountry,
					plantstate,
					possiblevalues,
					pretensioner,
					rearautomaticemergencybraking,
					rearcrosstrafficalert,
					rearvisibilitysystem,
					saeautomationlevel,
					saeautomationlevel_to,
					seatbeltsall,
					seatrows,
					seats,
					semiautomaticheadlampbeamswitching,
					series,
					series2,
					steeringlocation,
					suggestedvin,
					tpms,
					topspeedmph,
					trackwidth,
					tractioncontrol,
					trailerbodytype,
					trailerlength,
					trailertype,
					transmissionspeeds,
					transmissionstyle,
					trim,
					trim2,
					turbo,
					vin,
					valvetraindesign,
					vehicledescriptor,
					vehicletype,
					wheelbaselong,
					wheelbaseshort,
					wheelbasetype,
					wheelsizefront,
					wheelsizerear,
					wheels,
					windows
				)
				VALUES (
					$1,
					$2,
					$3,
					$4,
					$5,
					$6,
					$7,
					$8,
					$9,
					$10,
					$11,
					$12,
					$13,
					$14,
					$15,
					$16,
					$17,
					$18,
					$19,
					$20,
					$21,
					$22,
					$23,
					$24,
					$25,
					$26,
					$27,
					$28,
					$29,
					$30,
					$31,
					$32,
					$33,
					$34,
					$35,
					$36,
					$37,
					$38,
					$39,
					$40,
					$41,
					$42,
					$43,
					$44,
					$45,
					$46,
					$47,
					$48,
					$49,
					$50,
					$51,
					$52,
					$53,
					$54,
					$55,
					$56,
					$57,
					$58,
					$59,
					$60,
					$61,
					$62,
					$63,
					$64,
					$65,
					$66,
					$67,
					$68,
					$69,
					$70,
					$71,
					$72,
					$73,
					$74,
					$75,
					$76,
					$77,
					$78,
					$79,
					$80,
					$81,
					$82,
					$83,
					$84,
					$85,
					$86,
					$87,
					$88,
					$89,
					$90,
					$91,
					$92,
					$93,
					$94,
					$95,
					$96,
					$97,
					$98,
					$99,
					$100,
					$101,
					$102,
					$103,
					$104,
					$105,
					$106,
					$107,
					$108,
					$109,
					$110,
					$111,
					$112,
					$113,
					$114,
					$115,
					$116,
					$117,
					$118,
					$119,
					$120,
					$121,
					$122,
					$123,
					$124,
					$125,
					$126,
					$127,
					$128,
					$129,
					$130,
					$131,
					$132,
					$133,
					$134,
					$135,
					$136,
					$137,
					$138,
					$139,
					$140,
					$141,
					$142,
					$143,
					$144,
					$145,
					$146,
					$147,
					$148,
					$149,
					$150
				)
				RETURNING vin`
				v := ""
				err = db.QueryRow(s,
					d.ABS,
					d.ActiveSafetySysNote,
					d.AdaptiveCruiseControl,
					d.AdaptiveDrivingBeam,
					d.AdaptiveHeadlights,
					d.AdditionalErrorText,
					d.AirBagLocCurtain,
					d.AirBagLocFront,
					d.AirBagLocKnee,
					d.AirBagLocSeatCushion,
					d.AirBagLocSide,
					d.AutoReverseSystem,
					d.AutomaticPedestrianAlertingSound,
					d.AxleConfiguration,
					d.Axles,
					d.BasePrice,
					d.BatteryA,
					d.BatteryA_to,
					d.BatteryCells,
					d.BatteryInfo,
					d.BatteryKWh,
					d.BatteryKWh_to,
					d.BatteryModules,
					d.BatteryPacks,
					d.BatteryType,
					d.BatteryV,
					d.BatteryV_to,
					d.BedLengthIN,
					d.BedType,
					d.BlindSpotIntervention,
					d.BlindSpotMon,
					d.BodyCabType,
					d.BodyClass,
					d.BrakeSystemDesc,
					d.BrakeSystemType,
					d.BusFloorConfigType,
					d.BusLength,
					d.BusType,
					d.CAN_AACN,
					d.CIB,
					d.CashForClunkers,
					d.ChargerLevel,
					d.ChargerPowerKW,
					d.CoolingType,
					d.CurbWeightLB,
					d.CustomMotorcycleType,
					d.DaytimeRunningLight,
					d.DestinationMarket,
					d.DisplacementCC,
					d.DisplacementCI,
					d.DisplacementL,
					d.Doors,
					d.DriveType,
					d.DriverAssist,
					d.DynamicBrakeSupport,
					d.EDR,
					d.ESC,
					d.EVDriveUnit,
					d.ElectrificationLevel,
					d.EngineConfiguration,
					d.EngineCycles,
					d.EngineCylinders,
					d.EngineHP,
					d.EngineHP_to,
					d.EngineKW,
					d.EngineManufacturer,
					d.EngineModel,
					d.EntertainmentSystem,
					d.ErrorCode,
					d.ErrorText,
					d.ForwardCollisionWarning,
					d.FuelInjectionType,
					d.FuelTypePrimary,
					d.FuelTypeSecondary,
					d.GCWR,
					d.GCWR_to,
					d.GVWR,
					d.GVWR_to,
					d.KeylessIgnition,
					d.LaneCenteringAssistance,
					d.LaneDepartureWarning,
					d.LaneKeepSystem,
					d.LowerBeamHeadlampLightSource,
					d.Make,
					d.MakeID,
					d.Manufacturer,
					d.ManufacturerId,
					d.Model,
					d.ModelID,
					d.ModelYear,
					d.MotorcycleChassisType,
					d.MotorcycleSuspensionType,
					d.NCSABodyType,
					d.NCSAMake,
					d.NCSAMapExcApprovedBy,
					d.NCSAMapExcApprovedOn,
					d.NCSAMappingException,
					d.NCSAModel,
					d.NCSANote,
					d.NonLandUse,
					d.Note,
					d.OtherBusInfo,
					d.OtherEngineInfo,
					d.OtherMotorcycleInfo,
					d.OtherRestraintSystemInfo,
					d.OtherTrailerInfo,
					d.ParkAssist,
					d.PedestrianAutomaticEmergencyBraking,
					d.PlantCity,
					d.PlantCompanyName,
					d.PlantCountry,
					d.PlantState,
					d.PossibleValues,
					d.Pretensioner,
					d.RearAutomaticEmergencyBraking,
					d.RearCrossTrafficAlert,
					d.RearVisibilitySystem,
					d.SAEAutomationLevel,
					d.SAEAutomationLevel_to,
					d.SeatBeltsAll,
					d.SeatRows,
					d.Seats,
					d.SemiautomaticHeadlampBeamSwitching,
					d.Series,
					d.Series2,
					d.SteeringLocation,
					d.SuggestedVIN,
					d.TPMS,
					d.TopSpeedMPH,
					d.TrackWidth,
					d.TractionControl,
					d.TrailerBodyType,
					d.TrailerLength,
					d.TrailerType,
					d.TransmissionSpeeds,
					d.TransmissionStyle,
					d.Trim,
					d.Trim2,
					d.Turbo,
					d.VIN,
					d.ValveTrainDesign,
					d.VehicleDescriptor,
					d.VehicleType,
					d.WheelBaseLong,
					d.WheelBaseShort,
					d.WheelBaseType,
					d.WheelSizeFront,
					d.WheelSizeRear,
					d.Wheels,
					d.Windows,
				).Scan(&v)
				fmt.Print(err)
				if err != nil {
					errorResponse(&w, "An error occurred while attempting to insert the data into the database", err)
				}

				messageResponse(&w, fmt.Sprintf("Data for VIN %s was successfully inserted", v))
			}
		} else {
			// extract the vin from the query parameters
			v := q["vin"][0]
			fmt.Printf("  [%s] vin %s provided, will look up data to store\n", h, v)

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
	default:
		methodNotAllowedResponse(&w, h)
	}
}

// handle the request routes
func handleRequests() {
	http.HandleFunc("/", getBaseRoute)
	http.HandleFunc("/about", getAbout)
	http.HandleFunc("/vehicles", getVehicles)
	http.HandleFunc("/vinlookup", getVINLookup)
	http.HandleFunc("/storevindata", postStoreVINData)

	fmt.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// main
func main() {
	handleRequests()
}
