package model

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        *float64      `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type"`
}
