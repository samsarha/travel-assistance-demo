package datasource

import (
	"demo/back-end/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type OpenWeatherDataSource struct {
}

func (ow OpenWeatherDataSource) FetchWeatherData(
	countryCode string,
	countryDS CountryDataSource) [5]models.Weather {

	coordinates := countryDS.FetchCountryCoordinates(countryCode)

	var weather [5]models.Weather

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENWEATHERKEY")
	weatherResponse := MakeGetRequest(
		fmt.Sprintf(
			"https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&units=metric&appid=%s",
			coordinates.Latitude,
			coordinates.Longitude,
			apiKey))

	var opnWeatherResponseObj models.OpenWeather
	json.Unmarshal(weatherResponse, &opnWeatherResponseObj)

	for i := 0; i < 5; i++ {

		wthr := models.Weather{}

		wthr.Date = opnWeatherResponseObj.List[i].DtTxt
		wthr.Description = opnWeatherResponseObj.List[i].Weather[0].Description
		wthr.IconUri = opnWeatherResponseObj.List[i].Weather[0].Icon
		wthr.Max = opnWeatherResponseObj.List[i].Main.TempMax
		wthr.Min = opnWeatherResponseObj.List[i].Main.TempMin

		weather[i] = wthr
	}

	return weather

}
