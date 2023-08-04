package datasource

import "demo/back-end/models"

type ExchangeDataSource interface {
	FetchExchangeRate(countryCode string) models.ExchangeRate
}
