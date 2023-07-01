package exporter

import (
	"fmt"
	"strconv"
)

type kucoinExporter struct {
	next baseProvider
}

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
	rate, err := e.currentBTCPrice()
	if err == nil {
		return rate, nil
	}

	if e.next == nil {
		return 0, err
	}

	return e.next.CurrentBTCPrice()
}

func (e *kucoinExporter) SetNext(next baseProvider) {
	e.next = next
}

func (e kucoinExporter) currentBTCPrice() (float64, error) {
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
