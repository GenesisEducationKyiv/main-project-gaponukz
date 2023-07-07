package providers

import (
	"btcapp/src/entities"
	"fmt"
)

func NewCoinstatsProvider() baseProvider {
	return providerChainFactory("coinstats", func(from entities.Symbol, to entities.Symbol) (entities.Price, error) {
		word, err := fromSymbolToWord(from)
		if err != nil {
			return 0, err
		}

		apiUrl := fmt.Sprintf("https://api.coinstats.app/public/v1/coins/%s", word)
		var jsonResponse map[string]interface{}

		err = loadJsonResponseBody(fmt.Sprintf("%s?currency=%s", apiUrl, to), &jsonResponse)
		if err != nil {
			return 0, err
		}

		data, ok := jsonResponse["coin"].(map[string]interface{})
		if !ok {
			return 0, fmt.Errorf("data not found in the response")
		}

		price, ok := data["price"].(float64)
		if !ok {
			return 0, fmt.Errorf("data not found in the response")
		}

		return entities.Price(price), nil
	})
}
