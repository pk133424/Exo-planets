package handler

import (
	"encoding/json"
	"exo-planets/model"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type handler struct {
	data map[uuid.UUID]model.Exoplanet
}

type ExoplanetHandlerInterface interface {
	AddExoplanet(w http.ResponseWriter, r *http.Request)
	ListExoplanets(w http.ResponseWriter, r *http.Request)
	GetExoplanetByID(w http.ResponseWriter, r *http.Request)
	UpdateExoplanet(w http.ResponseWriter, r *http.Request)
	DeleteExoplanet(w http.ResponseWriter, r *http.Request)
	FuelEstimation(w http.ResponseWriter, r *http.Request)
}

func NewExoplanetHandler(data map[uuid.UUID]model.Exoplanet) ExoplanetHandlerInterface {
	return &handler{
		data: data,
	}
}

func (h handler) AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet model.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		log.Error().Err(err).Msg("AddExoplanet - json.NewDecoder")
		ErrorHandler(http.StatusBadRequest, "Bad Request", w)
		return
	}
	if !model.IsValidType(exoplanet.Type) {
		log.Error().Err(err).Msg("AddExoplanet - Invalid Type")
		ErrorHandler(http.StatusBadRequest, "Validation Error", w)
		return
	}
	if exoplanet.Type == model.Terrestrial && exoplanet.Mass == nil {
		log.Error().Err(err).Msg("AddExoplanet - Mass is required.")
		ErrorHandler(http.StatusBadRequest, "Validation Error", w)
		return
	}
	exoplanet.ID = uuid.New()
	err = validator.New().Struct(exoplanet)
	if err != nil {
		log.Error().Err(err).Msg("AddExoplanet - Validation Error")
		ErrorHandler(http.StatusBadRequest, "Validation Error", w)
		return
	}
	h.data[exoplanet.ID] = exoplanet
	resp, err := json.Marshal(exoplanet)
	if err != nil {
		log.Error().Err(err).Msg("AddExoplanet - error marshaling")
		ErrorHandler(http.StatusInternalServerError, "Internal server error", w)
	}
	SuccessResponseHandler(w, http.StatusOK, resp)
}

func (h handler) ListExoplanets(w http.ResponseWriter, r *http.Request) {
	planets := make([]model.Exoplanet, 0, len(h.data))
	for _, planet := range h.data {
		planets = append(planets, planet)
	}
	resp, err := json.Marshal(planets)
	if err != nil {
		log.Error().Err(err).Msg("ListExoplanets - error marshaling")
		ErrorHandler(http.StatusInternalServerError, "Internal server error", w)
	}
	SuccessResponseHandler(w, http.StatusOK, resp)
}

func (h handler) GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		ErrorHandler(http.StatusBadRequest, err.Error(), w)
		return
	}
	planet, exists := h.data[id]
	if !exists {
		ErrorHandler(http.StatusNotFound, "Exoplanet not found", w)
		return
	}
	resp, err := json.Marshal(planet)
	if err != nil {
		log.Error().Err(err).Msg("GetExoplanetByID - error marshaling")
		ErrorHandler(http.StatusInternalServerError, err.Error(), w)
	}
	SuccessResponseHandler(w, http.StatusOK, resp)
}

func (h handler) UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Error().Err(err).Msg("UpdateExoplanet - uuid.Parse")
		ErrorHandler(http.StatusBadRequest, err.Error(), w)
		return
	}
	_, exists := h.data[id]
	if !exists {
		ErrorHandler(http.StatusNotFound, "Exoplanet not found", w)
		return
	}
	var exoplanet model.Exoplanet
	err = json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		log.Error().Err(err).Msg("UpdateExoplanet - json.NewDecoder")
		ErrorHandler(http.StatusBadRequest, err.Error(), w)
		return
	}
	if !model.IsValidType(exoplanet.Type) {
		log.Error().Err(err).Msg("UpdateExoplanet - Invalid Type")
		ErrorHandler(http.StatusBadRequest, "Validation Error", w)
		return
	}
	if exoplanet.Type == model.Terrestrial && exoplanet.Mass == nil {
		log.Error().Err(err).Msg("UpdateExoplanet - Mass is required.")
		ErrorHandler(http.StatusBadRequest, "Validation Error", w)
		return
	}
	exoplanet.ID = id
	err = validator.New().Struct(exoplanet)
	if err != nil {
		log.Error().Err(err).Msg("UpdateExoplanet - Validation Error")
		ErrorHandler(http.StatusBadRequest, "Validation Error", w)
		return
	}
	h.data[id] = exoplanet
	resp, err := json.Marshal(exoplanet)
	if err != nil {
		log.Error().Err(err).Msg("UpdateExoplanet - error marshaling")
		ErrorHandler(http.StatusInternalServerError, err.Error(), w)
	}
	SuccessResponseHandler(w, http.StatusOK, resp)
}

func (h handler) DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Error().Err(err).Msg("DeleteExoplanet - uuid.Parse")
		ErrorHandler(http.StatusBadRequest, err.Error(), w)
		return
	}
	_, exists := h.data[id]
	if !exists {
		ErrorHandler(http.StatusNotFound, "Exoplanet not found", w)
		return
	}
	delete(h.data, id)
	SuccessResponseHandler(w, http.StatusOK, []byte("Deleted"))

}

func (h handler) FuelEstimation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Error().Err(err).Msg("FuelEstimation - uuid.Parse")
		ErrorHandler(http.StatusBadRequest, err.Error(), w)
		return
	}
	planet, exists := h.data[id]
	if !exists {
		log.Error().Err(err).Msg("FuelEstimation - Exoplanet not found")
		ErrorHandler(http.StatusNotFound, "Exoplanet not found", w)
		return
	}
	crewCapacity, err := strconv.Atoi(vars["crewCapacity"])
	if err != nil {
		log.Error().Err(err).Msg("FuelEstimation - strconv.Atoi")
		ErrorHandler(http.StatusBadRequest, err.Error(), w)
		return
	}

	var gravity float64
	switch planet.Type {
	case model.GasGiant:
		gravity = 0.5 / (planet.Radius * planet.Radius)
	case model.Terrestrial:
		if planet.Mass == nil {
			log.Error().Err(err).Msg("FuelEstimation - Mass is required for Terrestrial planets")
			ErrorHandler(http.StatusBadRequest, "Mass is required for Terrestrial planets", w)
			return
		}
		gravity = *planet.Mass / (planet.Radius * planet.Radius)
	}

	fuel := float64(planet.Distance) / (gravity * gravity) * float64(crewCapacity)
	resp, err := json.Marshal(map[string]float64{"fuel": fuel})
	if err != nil {
		log.Error().Err(err).Msg("FuelEstimation - error marshaling")
		ErrorHandler(http.StatusInternalServerError, err.Error(), w)
	}
	SuccessResponseHandler(w, http.StatusOK, resp)
}

func ErrorHandler(statusCode int, message string, w http.ResponseWriter) {
	body, _ := json.Marshal(&model.ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
	SuccessResponseHandler(w, statusCode, body)
}

func SuccessResponseHandler(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if statusCode != http.StatusNoContent {
		_, err := w.Write(body)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
}
