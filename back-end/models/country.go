package models

type Country struct {
	CountryName  string     `json:"countryname"`
	Gdp          float64    `json:"gdp"`
	Population   int        `json:"population"`
	ExchangeRate float64    `json:"exchangerate"`
	Weather      [5]Weather `json:"weather"`
}
