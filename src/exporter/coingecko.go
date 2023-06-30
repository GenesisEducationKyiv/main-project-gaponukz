package exporter

import (
	"fmt"
)

type coingeckoExporter struct{}

func NewCoingeckoExporter() *coingeckoExporter {
	return &coingeckoExporter{}
}

type coingeckoResponse struct {
	Bitcoin struct {
		UAH int `json:"uah"`
	} `json:"bitcoin"`
}

func (e coingeckoExporter) CurrentBTCPrice() (float64, error) {
	var apiResponse coingeckoResponse
	const ApiUrl = "https://api.coingecko.com/api/v3/simple/price"

	err := getJson(fmt.Sprintf("%s?ids=bitcoin&vs_currencies=uah", ApiUrl), &apiResponse)
	if err != nil {
		return 0, err
	}

	price := float64(apiResponse.Bitcoin.UAH)

	return price, nil
}
