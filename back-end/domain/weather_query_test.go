package domain

import (
	"demo/back-end/models"
	"testing"
)

type WeatherDataSourceStub struct {
	weather []models.Weather
}

func (ds WeatherDataSourceStub) FetchWeatherData(countryCode string) []models.Weather {
	return ds.weather
}

func TestQueryWeatherData(t *testing.T) {

	weatherArr := []models.Weather{{Date: "123454433443", IconUri: "dshdsdssdf", Min: 12, Max: 25, Description: "Rainy"}}

	datasource := WeatherDataSourceStub{weather: weatherArr}

	got := QueryWeatherData("MOZ", datasource)

	if len(got) != len(weatherArr) {
		t.Errorf("got %d, wanted %d", len(got), len(weatherArr))
	}
}
