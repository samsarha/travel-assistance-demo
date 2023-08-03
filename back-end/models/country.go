package models

type Country struct {
	CountryName  string    `json:countryname`
	Gdp          int       `json:gdp`
	ExchangeRate float64   `json:exchangerate`
	Weather      []Weather `json:weather`
}
