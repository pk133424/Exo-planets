package model

import "github.com/google/uuid"

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          uuid.UUID     `json:"id" validate:"required"`
	Name        string        `json:"name" validate:"required"`
	Description string        `json:"description" validate:"required"`
	Distance    int           `json:"distance" validate:"min=10,max=1000"`
	Radius      float64       `json:"radius" validate:"min=0.1,max=10"`
	Mass        *float64      `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type" validate:"required"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var validType = []ExoplanetType{GasGiant, Terrestrial}

func IsValidType(str ExoplanetType) bool {
	for _, v := range validType {
		if v == str {
			return true
		}
	}
	return false
}
