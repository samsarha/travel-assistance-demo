package main

import (
	"demo/back-end/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func countryByName(c *gin.Context) {
	//name := c.Param("name")
	weather := []models.Weather{{Date: "123454433443", IconUri: "dshdsdssdf", Min: 25, Max: 12, Description: "Rainy"}}
	country := models.Country{CountryName: "Zimbabwe", Gdp: 1000, ExchangeRate: 54.30, Weather: weather}
	c.IndentedJSON(http.StatusOK, country)

}

func main() {
	router := gin.Default()
	router.GET("/country/:name", countryByName)
	router.Run("localhost:8080")
}
