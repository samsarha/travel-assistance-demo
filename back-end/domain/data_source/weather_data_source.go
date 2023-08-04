package datasource

import "demo/back-end/models"

type WeatherDataSource interface {
	FetchWeatherData(countryCode string, countryDS CountryDataSource) [5]models.Weather
}
