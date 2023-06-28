package exporter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type coingeckoExporter struct{}

func NewCoingeckoExporter() *coingeckoExporter {
	return &coingeckoExporter{}
}

type currencyResponse struct {
	Bitcoin struct {
		UAH int `json:"uah"`
	} `json:"bitcoin"`
}

func (e coingeckoExporter) GetCurrentBTCPrice() (float64, error) {
	var apiResponse currencyResponse
	const ApiUrl = "https://api.coingecko.com/api/v3/simple/price"

	err := getJson(fmt.Sprintf("%s?ids=bitcoin&vs_currencies=uah", ApiUrl), &apiResponse)
	if err != nil {
		return 0, err
	}

	price := float64(apiResponse.Bitcoin.UAH)

	return price, nil
}

func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
