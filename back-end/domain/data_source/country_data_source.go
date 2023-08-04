package datasource

import "demo/back-end/models"

type CountryDataSource interface {
	FetchCountryData(countryCode string) models.Country
	FetchCountryCoordinates(countryCode string) models.Coordinates
}
