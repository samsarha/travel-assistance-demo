package datasource

import (
	"demo/back-end/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ExchgrteapiDataSource struct {
}

func (ex ExchgrteapiDataSource) FetchExchangeRate(countryCode string) models.ExchangeRate {

	exchangerate := models.ExchangeRate{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("EXCHANGERATEKEY")

	var currencyName = CountryCurr[strings.ToUpper(countryCode)]

	exchangeRateResponse := MakeGetRequest(
		fmt.Sprintf(
			"http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=%s",
			apiKey,
			currencyName))

	var exchangeRateResponseObj models.ExchangeRateAPI
	json.Unmarshal(exchangeRateResponse, &exchangeRateResponseObj)

	exchangerate.Rate = exchangeRateResponseObj.Rates[strings.ToUpper(currencyName)]

	return exchangerate

}
