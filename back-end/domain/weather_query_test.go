package domain

import (
	datasource "demo/back-end/domain/data_source"
	"demo/back-end/models"
	"testing"
)

type WeatherDataSourceStub struct {
	weather [5]models.Weather
}

func (ds WeatherDataSourceStub) FetchWeatherData(
	countryCode string,
	countryDS datasource.CountryDataSource) [5]models.Weather {

	return ds.weather
}

func TestQueryWeatherData(t *testing.T) {

	weatherArr := [5]models.Weather{{Date: "123454433443", IconUri: "dshdsdssdf", Min: 12, Max: 25, Description: "Rainy"}}

	weatherDS := WeatherDataSourceStub{weather: weatherArr}

	ctry := models.Country{CountryName: "Zimbabwe", Gdp: 1000, ExchangeRate: 54.30, Weather: weatherArr}

	countryDS := CountryDataSourceStub{country: ctry}

	got := QueryWeatherData("MOZ", weatherDS, countryDS)

	if len(got) != len(weatherArr) {
		t.Errorf("got %d, wanted %d", len(got), len(weatherArr))
	}
}
