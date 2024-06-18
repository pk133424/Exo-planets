package handler

import (
	"encoding/json"
	"errors"
	"exo-planets/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handler struct {
	data map[string]model.Exoplanet
}

type ExoplanetHandlerInterface interface {
	AddExoplanet(w http.ResponseWriter, r *http.Request)
	ListExoplanets(w http.ResponseWriter, r *http.Request)
	GetExoplanetByID(w http.ResponseWriter, r *http.Request)
	UpdateExoplanet(w http.ResponseWriter, r *http.Request)
	DeleteExoplanet(w http.ResponseWriter, r *http.Request)
	FuelEstimation(w http.ResponseWriter, r *http.Request)
}

func NewExoplanetHandler(data map[string]model.Exoplanet) ExoplanetHandlerInterface {
	return &handler{
		data: data,
	}
}

func (h handler) AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet model.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = validateExoplanet(exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.data[exoplanet.Name] = exoplanet
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func (h handler) ListExoplanets(w http.ResponseWriter, r *http.Request) {
	planets := make([]model.Exoplanet, 0, len(h.data))
	for _, planet := range h.data {
		planets = append(planets, planet)
	}
	json.NewEncoder(w).Encode(planets)
}

func (h handler) GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	planetName := mux.Vars(r)["name"]
	// id, err := uuid.Parse(vars["id"])
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	planet, exists := h.data[planetName]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(planet)
}

func (h handler) UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	planetName := mux.Vars(r)["name"]

	var exoplanet model.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.data[planetName] = exoplanet
	json.NewEncoder(w).Encode(exoplanet)
}

func (h handler) DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	planetName := mux.Vars(r)["name"]

	delete(h.data, planetName)
	w.WriteHeader(http.StatusNoContent)
}

func (h handler) FuelEstimation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	planet, exists := h.data[vars["name"]]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	crewCapacity, err := strconv.Atoi(vars["crewCapacity"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var gravity float64
	switch planet.Type {
	case model.GasGiant:
		gravity = 0.5 / (planet.Radius * planet.Radius)
	case model.Terrestrial:
		if planet.Mass == nil {
			http.Error(w, "Mass is required for Terrestrial planets", http.StatusBadRequest)
			return
		}
		gravity = *planet.Mass / (planet.Radius * planet.Radius)
	}

	fuel := float64(planet.Distance) / (gravity * gravity) * float64(crewCapacity)
	json.NewEncoder(w).Encode(map[string]float64{"fuel": fuel})
}

func validateExoplanet(exoplanet model.Exoplanet) error {
	if exoplanet.Distance <= 10 || exoplanet.Distance >= 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius <= 0.1 || exoplanet.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.Type == model.Terrestrial && (exoplanet.Mass == nil || *exoplanet.Mass <= 0.1 || *exoplanet.Mass >= 10) {
		return errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial planets")
	}
	return nil
}
