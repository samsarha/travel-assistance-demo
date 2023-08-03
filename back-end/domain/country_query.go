package domain

import (
	datasource "demo/back-end/domain/data_source"
	"demo/back-end/models"
)

func QueryCountryData(countryCode string, datasource datasource.CountryDataSource) models.Country {
	return datasource.FetchCountryData(countryCode)
}
