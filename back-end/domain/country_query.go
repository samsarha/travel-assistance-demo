package domain

import (
	datasource "demo/back-end/domain/data_source"
	"demo/back-end/models"
)

func QueryCountryData(countryCode string, datasource datasource.CountryDataSource) models.Country {

	weather := []models.Weather{{Date: "123454433443", IconUri: "dshdsdssdf", Min: 25, Max: 12, Description: "Rainy"}}
	country := models.Country{CountryName: "Zimbabwe", Gdp: 1000, ExchangeRate: 54.30, Weather: weather}
	return country
}
