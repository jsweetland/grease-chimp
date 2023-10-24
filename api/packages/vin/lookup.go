package vin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Lookup(vin string) (data LookupResponse, err error) {
	base := "https://vpic.nhtsa.dot.gov/api/vehicles/DecodeVinValues"
	params := "?format=json"
	u := fmt.Sprintf("%s/%s%s", base, vin, params)

	r, err := http.Get(u)
	if err != nil {
		return EmptyResponse, err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return EmptyResponse, err
	}

	s := string(b)
	d := LookupResponse{}

	err = json.Unmarshal([]byte(s), &d)
	if err != nil {
		return EmptyResponse, err
	}

	return d, nil
}
