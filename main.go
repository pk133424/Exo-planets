package main

import (
	"exo-planets/handler"
	"exo-planets/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	var exoplanets = make(map[string]model.Exoplanet)

	exoplanetHandler := handler.NewExoplanetHandler(exoplanets)
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", exoplanetHandler.AddExoplanet).Methods("POST")
	r.HandleFunc("/exoplanets", exoplanetHandler.ListExoplanets).Methods("GET")
	r.HandleFunc("/exoplanets/{name}", exoplanetHandler.GetExoplanetByID).Methods("GET")
	r.HandleFunc("/exoplanets/{name}", exoplanetHandler.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/exoplanets/{name}", exoplanetHandler.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/exoplanets/{name}/fuel/{crewCapacity}", exoplanetHandler.FuelEstimation).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().Err(err).Msg("error on http.ListenAndServe")
	}

}
