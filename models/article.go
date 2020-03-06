package models


type Article struct {
	ID          string `json:"Id"`
	Description string `json:"Description"`
	VatCode     int    `json:"VatCode"`
	Price       string `json:"Price"`
	Unit        string `json:"Unit"`
}
