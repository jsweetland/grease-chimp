module main

go 1.21.0

require (
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/gc/db v0.0.0
	github.com/gc/vin v0.0.0
	github.com/gc/handlers v0.0.0
	github.com/gc/types v0.0.0
	github.com/gc/helpers v0.0.0
	github.com/gc/emptydata v0.0.0
)

replace (
	"github.com/gc/db" v0.0.0 => "./packages/db"
	"github.com/gc/vin" v0.0.0 => "./packages/vin"
	"github.com/gc/handlers" v0.0.0 => "./packages/handlers"
	"github.com/gc/types" v0.0.0 => "./packges/types"
	"github.com/gc/helpers" v0.0.0 => "./packages/helpers"
	"github.com/gc/emptydata" v0.0.0 => "./packages/emptydata"
)
