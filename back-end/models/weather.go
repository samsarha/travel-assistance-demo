package models

type Weather struct {
	Date        string `json:"countryname"`
	IconUri     string `json:"gdp"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Description string `json:"description"`
}
