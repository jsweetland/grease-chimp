package types

// type for the database config
type DBConf struct {
	User string
	Pass string
	Name string
	Host string
}

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

// type for the vehicle license plate
type Plate struct {
	Issuer string `json:"issuer,omitempty"`
	Value  string `json:"value,omitempty"`
}

// type for the vehicle data
type Vehicle struct {
	ID         int    `json:"id"`
	VIN        string `json:"vin,omitempty"`
	Make       string `json:"make,omitempty"`
	Plate      Plate  `json:"plate,omitempty"`
	PlateValue string `json:"platevalue,omitempty"`
	Model      string `json:"model,omitempty"`
	Year       int    `json:"year,omitempty"`
	Trim       string `json:"trim,omitempty"`
	Package    string `json:"package,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	Color      Color  `json:"color,omitempty"`
}

type VINData struct {
	ABS                                 string
	ActiveSafetySysNote                 string
	AdaptiveCruiseControl               string
	AdaptiveDrivingBeam                 string
	AdaptiveHeadlights                  string
	AdditionalErrorText                 string
	AirBagLocCurtain                    string
	AirBagLocFront                      string
	AirBagLocKnee                       string
	AirBagLocSeatCushion                string
	AirBagLocSide                       string
	AutoReverseSystem                   string
	AutomaticPedestrianAlertingSound    string
	AxleConfiguration                   string
	Axles                               string
	BasePrice                           string
	BatteryA                            string
	BatteryA_to                         string
	BatteryCells                        string
	BatteryInfo                         string
	BatteryKWh                          string
	BatteryKWh_to                       string
	BatteryModules                      string
	BatteryPacks                        string
	BatteryType                         string
	BatteryV                            string
	BatteryV_to                         string
	BedLengthIN                         string
	BedType                             string
	BlindSpotIntervention               string
	BlindSpotMon                        string
	BodyCabType                         string
	BodyClass                           string
	BrakeSystemDesc                     string
	BrakeSystemType                     string
	BusFloorConfigType                  string
	BusLength                           string
	BusType                             string
	CAN_AACN                            string
	CIB                                 string
	CashForClunkers                     string
	ChargerLevel                        string
	ChargerPowerKW                      string
	CoolingType                         string
	CurbWeightLB                        string
	CustomMotorcycleType                string
	DaytimeRunningLight                 string
	DestinationMarket                   string
	DisplacementCC                      string
	DisplacementCI                      string
	DisplacementL                       string
	Doors                               string
	DriveType                           string
	DriverAssist                        string
	DynamicBrakeSupport                 string
	EDR                                 string
	ESC                                 string
	EVDriveUnit                         string
	ElectrificationLevel                string
	EngineConfiguration                 string
	EngineCycles                        string
	EngineCylinders                     string
	EngineHP                            string
	EngineHP_to                         string
	EngineKW                            string
	EngineManufacturer                  string
	EngineModel                         string
	EntertainmentSystem                 string
	ErrorCode                           string
	ErrorText                           string
	ForwardCollisionWarning             string
	FuelInjectionType                   string
	FuelTypePrimary                     string
	FuelTypeSecondary                   string
	GCWR                                string
	GCWR_to                             string
	GVWR                                string
	GVWR_to                             string
	KeylessIgnition                     string
	LaneCenteringAssistance             string
	LaneDepartureWarning                string
	LaneKeepSystem                      string
	LowerBeamHeadlampLightSource        string
	Make                                string
	MakeID                              string
	Manufacturer                        string
	ManufacturerId                      string
	Model                               string
	ModelID                             string
	ModelYear                           string
	MotorcycleChassisType               string
	MotorcycleSuspensionType            string
	NCSABodyType                        string
	NCSAMake                            string
	NCSAMapExcApprovedBy                string
	NCSAMapExcApprovedOn                string
	NCSAMappingException                string
	NCSAModel                           string
	NCSANote                            string
	NonLandUse                          string
	Note                                string
	OtherBusInfo                        string
	OtherEngineInfo                     string
	OtherMotorcycleInfo                 string
	OtherRestraintSystemInfo            string
	OtherTrailerInfo                    string
	ParkAssist                          string
	PedestrianAutomaticEmergencyBraking string
	PlantCity                           string
	PlantCompanyName                    string
	PlantCountry                        string
	PlantState                          string
	PossibleValues                      string
	Pretensioner                        string
	RearAutomaticEmergencyBraking       string
	RearCrossTrafficAlert               string
	RearVisibilitySystem                string
	SAEAutomationLevel                  string
	SAEAutomationLevel_to               string
	SeatBeltsAll                        string
	SeatRows                            string
	Seats                               string
	SemiautomaticHeadlampBeamSwitching  string
	Series                              string
	Series2                             string
	SteeringLocation                    string
	SuggestedVIN                        string
	TPMS                                string
	TopSpeedMPH                         string
	TrackWidth                          string
	TractionControl                     string
	TrailerBodyType                     string
	TrailerLength                       string
	TrailerType                         string
	TransmissionSpeeds                  string
	TransmissionStyle                   string
	Trim                                string
	Trim2                               string
	Turbo                               string
	VIN                                 string
	ValveTrainDesign                    string
	VehicleDescriptor                   string
	VehicleType                         string
	WheelBaseLong                       string
	WheelBaseShort                      string
	WheelBaseType                       string
	WheelSizeFront                      string
	WheelSizeRear                       string
	Wheels                              string
	Windows                             string
}

type VINLookupResponse struct {
	Count          int
	Message        string
	SearchCriteria string
	Results        []VINData
}
