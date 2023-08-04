package datasource

import (
	"demo/back-end/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type WorldBankDataSource struct {
}

func (wb WorldBankDataSource) FetchCountryData(countryCode string) models.Country {

	country := models.Country{}

	populationResponse := MakeGetRequest("https://api.worldbank.org/v2/country/moz/indicator/SP.POP.TOTL?date=2022&format=json")

	var populationResponseObj models.Population
	json.Unmarshal(populationResponse, &populationResponseObj)

	country.CountryName = populationResponseObj[1][0].Country.Value
	country.Population = populationResponseObj[1][0].Value

	gdpResponse := MakeGetRequest("https://api.worldbank.org/v2/country/moz/indicator/NY.GDP.MKTP.CD?date=2022&format=json")

	var gdpResponseObj models.Gdp
	json.Unmarshal(gdpResponse, &gdpResponseObj)

	country.Gdp = gdpResponseObj[1][0].Value

	return country
}

func (wb WorldBankDataSource) FetchCountryCoordinates(countryCode string) models.Coordinates {

	coordinates := models.Coordinates{}

	cityResponse := MakeGetRequest("https://api.worldbank.org/v2/country/mz?format=json")

	var cityResponseObj models.CapitalCity
	json.Unmarshal(cityResponse, &cityResponseObj)

	coordinates.Latitude = cityResponseObj[1][0].Latitude
	coordinates.Longitude = cityResponseObj[1][0].Longitude

	return coordinates
}

func MakeGetRequest(endpoint string) []byte {

	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	return responseData
}
