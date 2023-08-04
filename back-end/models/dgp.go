package models

type Gdp [][]struct {
	Indicator struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"indicator"`

	Country struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"country"`

	Countryiso3Code string  `json:"countryiso3code"`
	Date            string  `json:"date"`
	Value           float64 `json:"value"`
	Unit            string  `json:"unit"`
	ObsStatus       string  `json:"obs_status"`
	Decimal         int     `json:"decimal"`
}
