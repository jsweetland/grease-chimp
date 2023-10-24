package types

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
	ID       int    `json:"id"`
	VIN      string `json:"vin,omitempty"`
	Make     string `json:"make,omitempty"`
	Model    string `json:"model,omitempty"`
	Year     int    `json:"year,omitempty"`
	Trim     string `json:"trim,omitempty"`
	Package  string `json:"package,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Color    Color  `json:"color,omitempty"`
}
