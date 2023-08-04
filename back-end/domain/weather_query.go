package domain

import (
	datasource "demo/back-end/domain/data_source"
	"demo/back-end/models"
)

func QueryWeatherData(
	countryCode string,
	datasource datasource.WeatherDataSource) []models.Weather {
	return datasource.FetchWeatherData(countryCode)
}
