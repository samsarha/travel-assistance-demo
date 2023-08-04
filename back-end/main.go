package main

import (
	"demo/back-end/domain"
	datasource "demo/back-end/domain/data_source"
	"demo/back-end/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func countryByName(c *gin.Context) {
	countryCode := c.Param("code")

	wbankDS := datasource.WorldBankDataSource{}
	country := domain.QueryCountryData(countryCode, wbankDS)
	weather := []models.Weather{{Date: "123454433443", IconUri: "dshdsdssdf", Min: 25, Max: 12, Description: "Rainy"}}
	country.Weather = weather
	c.IndentedJSON(http.StatusOK, country)

}

func main() {
	router := gin.Default()
	router.GET("/country/:code", countryByName)
	router.Run("localhost:8080")
}
