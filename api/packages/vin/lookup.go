package vin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gc/emptydata"
	"github.com/gc/types"
)

func Lookup(vin string) (data types.VINLookupResponse, err error) {
	base := "https://vpic.nhtsa.dot.gov/api/vehicles/DecodeVinValues"
	params := "?format=json"
	u := fmt.Sprintf("%s/%s%s", base, vin, params)

	r, err := http.Get(u)
	if err != nil {
		return emptydata.EmptyVINLookupResponse, err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return emptydata.EmptyVINLookupResponse, err
	}

	s := string(b)
	d := types.VINLookupResponse{}

	err = json.Unmarshal([]byte(s), &d)
	if err != nil {
		return emptydata.EmptyVINLookupResponse, err
	}

	return d, nil
}
