package models

type Population [][]struct {
	Indicator struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"indicator"`
	Country struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"country"`
	Countryiso3Code string `json:"countryiso3code"`
	Date            string `json:"date"`
	Value           int    `json:"value"`
	Unit            string `json:"unit"`
	ObsStatus       string `json:"obs_status"`
	Decimal         int    `json:"decimal"`
}
