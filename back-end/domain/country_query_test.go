package domain

import (
	"demo/back-end/models"
	"testing"
)

type CountryDataSourceStub struct {
	country models.Country
}

func (ds CountryDataSourceStub) FetchCountryData(countryCode string) models.Country {
	return ds.country
}

func (ds CountryDataSourceStub) FetchCountryCoordinates(countryCode string) models.Coordinates {
	return models.Coordinates{Latitude: "1233", Longitude: "-1233.2"}
}

func TestQueryCountryData(t *testing.T) {

	weather := [5]models.Weather{{Date: "123454433443", IconUri: "dshdsdssdf", Min: 25, Max: 12, Description: "Rainy"}}
	want := models.Country{CountryName: "Zimbabwe", Gdp: 1000, ExchangeRate: 54.30, Weather: weather}

	datasource := CountryDataSourceStub{country: want}

	got := QueryCountryData("MOZ", datasource)

	if got.CountryName != want.CountryName {
		t.Errorf("got %q, wanted %q", got.CountryName, want.CountryName)
	}
}

func TestQueryCountryDataWeather(t *testing.T) {

	wthr := [5]models.Weather{{Date: "123454433443", IconUri: "dshdsdssdf", Min: 25, Max: 12, Description: "Rainy"}}
	cntry := models.Country{CountryName: "Zimbabwe", Gdp: 1000, ExchangeRate: 54.30, Weather: wthr}

	datasource := CountryDataSourceStub{country: cntry}

	got := len(QueryCountryData("MZ", datasource).Weather)

	if got < 1 {
		t.Errorf("got %q, wanted at least one weather forecast", got)
	}
}
