package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gc/handlers"
)

// -----   Route Handlers   -----

func startServer() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handlers.GetBaseRoute)
	r.Get("/about", handlers.GetAbout)
	r.Get("/vehicles", handlers.GetVehicles)
	r.Get("/vehicle/{id}", handlers.GetVehicleByID)
	r.Post("/vehicle", handlers.PostAddOrUpdateVehicle)
	r.Get("/vinlookup", handlers.GetVINLookup)
	r.Post("/storevindata", handlers.PostStoreVINData)

	p := "10000"
	fmt.Printf("Listening on port %s\n", p)
	http.ListenAndServe(fmt.Sprintf(":%s", p), r)
}

func main() {
	startServer()
}
