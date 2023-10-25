package emptydata

import (
	"github.com/gc/types"
)

var EmptyVINDataList []types.VINData

var EmptyVINLookupResponse = types.VINLookupResponse{
	Count:          0,
	Message:        "",
	SearchCriteria: "",
	Results:        EmptyVINDataList,
}
