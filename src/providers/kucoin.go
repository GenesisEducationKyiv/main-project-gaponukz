package providers

import (
	"btcapp/src/entities"
	"fmt"
	"strconv"
)

func NewKucoinProvider() baseProvider {
	return providerChainFactory("kucoin", func(from entities.Symbol, to entities.Symbol) (entities.Price, error) {
		const apiUrl = "https://api.kucoin.com/api/v1/prices"
		url := fmt.Sprintf("%s?currencies=%s&base=%s", apiUrl, from, to)

		var jsonResponse map[string]interface{}
		err := loadJsonResponseBody(url, &jsonResponse)
		if err != nil {
			return 0, err
		}

		data, ok := jsonResponse["data"].(map[string]interface{})
		if !ok {
			return 0, fmt.Errorf("data not found in the response")
		}

		rateString, ok := data[string(from)].(string)
		if !ok {
			return 0, fmt.Errorf("rate for symbol %s not found", from)
		}

		rate, err := strconv.ParseFloat(rateString, 64)
		if err != nil {
			return 0, err
		}

		return entities.Price(rate), nil
	})
}
