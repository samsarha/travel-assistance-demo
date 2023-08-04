package domain

import (
	datasource "demo/back-end/domain/data_source"
	"demo/back-end/models"
)

func QueryExchangeRateData(countryCode string, datasource datasource.ExchangeDataSource) models.ExchangeRate {
	return datasource.FetchExchangeRate(countryCode)
}
