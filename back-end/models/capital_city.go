package models

type CapitalCity [][]struct {
	ID       string `json:"id"`
	Iso2Code string `json:"iso2Code"`
	Name     string `json:"name"`
	Region   struct {
		ID       string `json:"id"`
		Iso2Code string `json:"iso2code"`
		Value    string `json:"value"`
	} `json:"region"`
	Adminregion struct {
		ID       string `json:"id"`
		Iso2Code string `json:"iso2code"`
		Value    string `json:"value"`
	} `json:"adminregion"`
	IncomeLevel struct {
		ID       string `json:"id"`
		Iso2Code string `json:"iso2code"`
		Value    string `json:"value"`
	} `json:"incomeLevel"`
	LendingType struct {
		ID       string `json:"id"`
		Iso2Code string `json:"iso2code"`
		Value    string `json:"value"`
	} `json:"lendingType"`
	CapitalCity string `json:"capitalCity"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
}
