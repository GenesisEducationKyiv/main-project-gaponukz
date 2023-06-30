package exporter

import (
	"fmt"
	"strconv"
)

type kucoinExporter struct{}

func NewKucoinExporter() kucoinExporter {
	return kucoinExporter{}
}

type kucoinResponse struct {
	Code string `json:"code"`
	Data struct {
		BTC string `json:"BTC"`
	} `json:"data"`
}

func (e kucoinExporter) CurrentBTCPrice() (float64, error) {
	var apiResponse kucoinResponse
	const ApiUrl = "https://api.kucoin.com/api/v1/prices"

	err := getJson(fmt.Sprintf("%s?currencies=BTC&base=UAH", ApiUrl), &apiResponse)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(apiResponse.Data.BTC, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
