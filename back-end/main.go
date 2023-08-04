package main

import (
	"demo/back-end/domain"
	datasource "demo/back-end/domain/data_source"
	"net/http"

	"github.com/gin-gonic/gin"
)

func countryByName(c *gin.Context) {
	countryCode := c.Param("code")

	wBankDS := datasource.WorldBankDataSource{}
	country := domain.QueryCountryData(countryCode, wBankDS)

	oWeatherDS := datasource.OpenWeatherDataSource{}
	weather := domain.QueryWeatherData(countryCode, oWeatherDS, wBankDS)
	country.Weather = weather

	exchangeRateApiDS := datasource.ExchgrteapiDataSource{}
	exchangeRate := domain.QueryExchangeRateData(countryCode, exchangeRateApiDS)
	country.ExchangeRate = exchangeRate.Rate

	c.IndentedJSON(http.StatusOK, country)

}

func main() {
	router := gin.Default()
	router.GET("/country/:code", countryByName)
	router.Run("localhost:8080")
}
