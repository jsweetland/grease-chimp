module main

go 1.21.0

require (
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/gc/db v0.0.0
	github.com/gc/vin v0.0.0
)

replace (
	"github.com/gc/db" v0.0.0 => "./packages/db"
	"github.com/gc/vin" v0.0.0 => "./packages/vin"
)
