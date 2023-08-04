package datasource

import "demo/back-end/models"

type WeatherDataSource interface {
	FetchWeatherData(countryCode string) []models.Weather
}
