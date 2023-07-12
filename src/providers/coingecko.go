package providers

import (
	"btcapp/src/entities"
	"fmt"
	"strings"
)

func NewCoingeckoProvider() baseProvider {
	return providerChainFactory("coingecko", func(from entities.Symbol, to entities.Symbol) (entities.Price, error) {
		const apiUrl = "https://api.coingecko.com/api/v3/simple/price"
		lowSymbol := strings.ToLower(string(to))

		word, err := fromSymbolToWord(from)
		if err != nil {
			return 0, err
		}

		var jsonResponse map[string]interface{}
		err = loadJsonResponseBody(fmt.Sprintf("%s?ids=%s&vs_currencies=%s", apiUrl, word, lowSymbol), &jsonResponse)
		if err != nil {
			return 0, err
		}

		data, ok := jsonResponse[word].(map[string]interface{})
		if !ok {
			return 0, fmt.Errorf("data not found in the response")
		}

		price, ok := data[lowSymbol].(float64)
		if !ok {
			return 0, fmt.Errorf("data not found in the response")
		}

		return entities.Price(price), nil
	})
}
