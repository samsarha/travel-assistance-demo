package domain

import (
	datasource "demo/back-end/domain/data_source"
	"demo/back-end/models"
)

func QueryWeatherData(
	countryCode string,
	weatherDS datasource.WeatherDataSource,
	countryDS datasource.CountryDataSource) [5]models.Weather {

	return weatherDS.FetchWeatherData(countryCode, countryDS)
}
