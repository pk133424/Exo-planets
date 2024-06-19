package main

import (
	"exo-planets/handler"
	"exo-planets/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	earthMass := 1.0
	earthId := uuid.New()
	exoplanets := map[uuid.UUID]model.Exoplanet{
		earthId: {
			ID:          earthId,
			Name:        "Earth",
			Description: "Our home planet.",
			Distance:    0,
			Radius:      1.0,
			Mass:        &earthMass,
			Type:        "Terrestrial",
		},
	}

	log.Info().Str("id", earthId.String()).Msg("Earth planet auto created")

	exoplanetHandler := handler.NewExoplanetHandler(exoplanets)
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", exoplanetHandler.AddExoplanet).Methods("POST")
	r.HandleFunc("/exoplanets", exoplanetHandler.ListExoplanets).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", exoplanetHandler.GetExoplanetByID).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", exoplanetHandler.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", exoplanetHandler.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel/{crewCapacity}", exoplanetHandler.FuelEstimation).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().Err(err).Msg("error on http.ListenAndServe")
	}

}
