package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gc/types"

	_ "github.com/lib/pq"
)

type DBConf struct {
	User string
	Pass string
	Name string
	Host string
}

var dbConfig = DBConf{
	User: "gcuser",
	Pass: "gcpass",
	Name: "gc",
	Host: "localhost",
}

// generate the database connection string
func GenerateConnStr(c DBConf) string {
	s := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Name,
	)

	return s
}

// connect to the database
func Connect() (db *sql.DB, err error) {
	// generate the database connection string
	s := GenerateConnStr(dbConfig)

	// connect to the database
	db, err = sql.Open("postgres", s)
	if err != nil {
		return nil, err
	}

	// return the database connection
	return db, nil
}

// insert vin data
func InsertVINData(db *sql.DB, d types.VINData) (vin string, err error) {
	v := ""

	s := `
	INSERT INTO vindata (abs, activesafetysysnote, adaptivecruisecontrol, adaptivedrivingbeam, adaptiveheadlights, additionalerrortext, airbagloccurtain, airbaglocfront, airbaglocknee, airbaglocseatcushion, airbaglocside, autoreversesystem, automaticpedestrianalertingsound, axleconfiguration, axles, baseprice, batterya, batterya_to, batterycells, batteryinfo, batterykwh, batterykwh_to, batterymodules, batterypacks, batterytype, batteryv, batteryv_to, bedlengthin, bedtype, blindspotintervention, blindspotmon, bodycabtype, bodyclass, brakesystemdesc, brakesystemtype, busfloorconfigtype, buslength, bustype, can_aacn, cib, cashforclunkers, chargerlevel, chargerpowerkw, coolingtype, curbweightlb, custommotorcycletype, daytimerunninglight, destinationmarket, displacementcc, displacementci, displacementl, doors, drivetype, driverassist, dynamicbrakesupport, edr, esc, evdriveunit, electrificationlevel, engineconfiguration, enginecycles, enginecylinders, enginehp, enginehp_to, enginekw, enginemanufacturer, enginemodel, entertainmentsystem, errorcode, errortext, forwardcollisionwarning, fuelinjectiontype, fueltypeprimary, fueltypesecondary, gcwr, gcwr_to, gvwr, gvwr_to, keylessignition, lanecenteringassistance, lanedeparturewarning, lanekeepsystem, lowerbeamheadlamplightsource, make, makeid, manufacturer, manufacturerid, model, modelid, modelyear, motorcyclechassistype, motorcyclesuspensiontype, ncsabodytype, ncsamake, ncsamapexcapprovedby, ncsamapexcapprovedon, ncsamappingexception, ncsamodel, ncsanote, nonlanduse, note, otherbusinfo, otherengineinfo, othermotorcycleinfo, otherrestraintsysteminfo, othertrailerinfo, parkassist, pedestrianautomaticemergencybraking, plantcity, plantcompanyname, plantcountry, plantstate, possiblevalues, pretensioner, rearautomaticemergencybraking, rearcrosstrafficalert, rearvisibilitysystem, saeautomationlevel, saeautomationlevel_to, seatbeltsall, seatrows, seats, semiautomaticheadlampbeamswitching, series, series2, steeringlocation, suggestedvin, tpms, topspeedmph, trackwidth, tractioncontrol, trailerbodytype, trailerlength, trailertype, transmissionspeeds, transmissionstyle, trim, trim2, turbo, vin, valvetraindesign, vehicledescriptor, vehicletype, wheelbaselong, wheelbaseshort, wheelbasetype, wheelsizefront, wheelsizerear, wheels, windows)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81, $82, $83, $84, $85, $86, $87, $88, $89, $90, $91, $92, $93, $94, $95, $96, $97, $98, $99, $100, $101, $102, $103, $104, $105, $106, $107, $108, $109, $110, $111, $112, $113, $114, $115, $116, $117, $118, $119, $120, $121, $122, $123, $124, $125, $126, $127, $128, $129, $130, $131, $132, $133, $134, $135, $136, $137, $138, $139, $140, $141, $142, $143, $144, $145, $146, $147, $148, $149, $150)
	RETURNING vin`
	err = db.QueryRow(s, d.ABS, d.ActiveSafetySysNote, d.AdaptiveCruiseControl, d.AdaptiveDrivingBeam, d.AdaptiveHeadlights, d.AdditionalErrorText, d.AirBagLocCurtain, d.AirBagLocFront, d.AirBagLocKnee, d.AirBagLocSeatCushion, d.AirBagLocSide, d.AutoReverseSystem, d.AutomaticPedestrianAlertingSound, d.AxleConfiguration, d.Axles, d.BasePrice, d.BatteryA, d.BatteryA_to, d.BatteryCells, d.BatteryInfo, d.BatteryKWh, d.BatteryKWh_to, d.BatteryModules, d.BatteryPacks, d.BatteryType, d.BatteryV, d.BatteryV_to, d.BedLengthIN, d.BedType, d.BlindSpotIntervention, d.BlindSpotMon, d.BodyCabType, d.BodyClass, d.BrakeSystemDesc, d.BrakeSystemType, d.BusFloorConfigType, d.BusLength, d.BusType, d.CAN_AACN, d.CIB, d.CashForClunkers, d.ChargerLevel, d.ChargerPowerKW, d.CoolingType, d.CurbWeightLB, d.CustomMotorcycleType, d.DaytimeRunningLight, d.DestinationMarket, d.DisplacementCC, d.DisplacementCI, d.DisplacementL, d.Doors, d.DriveType, d.DriverAssist, d.DynamicBrakeSupport, d.EDR, d.ESC, d.EVDriveUnit, d.ElectrificationLevel, d.EngineConfiguration, d.EngineCycles, d.EngineCylinders, d.EngineHP, d.EngineHP_to, d.EngineKW, d.EngineManufacturer, d.EngineModel, d.EntertainmentSystem, d.ErrorCode, d.ErrorText, d.ForwardCollisionWarning, d.FuelInjectionType, d.FuelTypePrimary, d.FuelTypeSecondary, d.GCWR, d.GCWR_to, d.GVWR, d.GVWR_to, d.KeylessIgnition, d.LaneCenteringAssistance, d.LaneDepartureWarning, d.LaneKeepSystem, d.LowerBeamHeadlampLightSource, d.Make, d.MakeID, d.Manufacturer, d.ManufacturerId, d.Model, d.ModelID, d.ModelYear, d.MotorcycleChassisType, d.MotorcycleSuspensionType, d.NCSABodyType, d.NCSAMake, d.NCSAMapExcApprovedBy, d.NCSAMapExcApprovedOn, d.NCSAMappingException, d.NCSAModel, d.NCSANote, d.NonLandUse, d.Note, d.OtherBusInfo, d.OtherEngineInfo, d.OtherMotorcycleInfo, d.OtherRestraintSystemInfo, d.OtherTrailerInfo, d.ParkAssist, d.PedestrianAutomaticEmergencyBraking, d.PlantCity, d.PlantCompanyName, d.PlantCountry, d.PlantState, d.PossibleValues, d.Pretensioner, d.RearAutomaticEmergencyBraking, d.RearCrossTrafficAlert, d.RearVisibilitySystem, d.SAEAutomationLevel, d.SAEAutomationLevel_to, d.SeatBeltsAll, d.SeatRows, d.Seats, d.SemiautomaticHeadlampBeamSwitching, d.Series, d.Series2, d.SteeringLocation, d.SuggestedVIN, d.TPMS, d.TopSpeedMPH, d.TrackWidth, d.TractionControl, d.TrailerBodyType, d.TrailerLength, d.TrailerType, d.TransmissionSpeeds, d.TransmissionStyle, d.Trim, d.Trim2, d.Turbo, d.VIN, d.ValveTrainDesign, d.VehicleDescriptor, d.VehicleType, d.WheelBaseLong, d.WheelBaseShort, d.WheelBaseType, d.WheelSizeFront, d.WheelSizeRear, d.Wheels, d.Windows).Scan(&v)
	if err != nil {
		return "", err
	}

	return v, nil
}

// insert new vehicle data
func InsertVehicleData(db *sql.DB, v types.Vehicle) (id int, err error) {
	id = 0

	s := `
	INSERT INTO vehicles (vin, make, model, year, trim, package, nickname, colorname, colorhex)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id`
	err = db.QueryRow(s, v.VIN, v.Make, v.Model, v.Year, v.Trim, v.Package, v.Nickname, v.Color.Hex, v.Color.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// update existing vehicle data
func UpdateVehicleData(db *sql.DB, v types.Vehicle) (err error) {
	s := `
	UPDATE vehicles
	SET vin = $2, make = $3, model = $4, year = $5, trim = $6, package = $7, nickname = $8, colorname = $9, colorhex = $10
	WHERE id = $1`
	res, err := db.Exec(s, v.ID, v.VIN, v.Make, v.Model, v.Year, v.Trim, v.Package, v.Nickname, v.Color.Hex, v.Color.Name)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("no rows were affected when attempting to update a vehicle")
	} else if count > 1 {
		return errors.New("more than one row was affected when attempting to update a single vehicle")
	}

	return nil
}
