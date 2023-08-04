package models

type Weather struct {
	Date        string  `json:"countryname"`
	IconUri     string  `json:"gdp"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
	Description string  `json:"description"`
}
