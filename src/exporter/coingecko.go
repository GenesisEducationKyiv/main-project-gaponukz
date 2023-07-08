package exporter

import (
	"fmt"
)

type coingeckoResponse struct {
	Bitcoin struct {
		UAH int `json:"uah"`
	} `json:"bitcoin"`
}

func NewCoingeckoExporter() baseProvider {
	return providerChainFactory("coingecko", func() (float64, error) {
		var apiResponse coingeckoResponse
		const ApiUrl = "https://api.coingecko.com/api/v3/simple/price"

		err := loadJsonResponseBody(fmt.Sprintf("%s?ids=bitcoin&vs_currencies=uah", ApiUrl), &apiResponse)
		if err != nil {
			return 0, err
		}

		price := float64(apiResponse.Bitcoin.UAH)

		return price, nil
	})
}
