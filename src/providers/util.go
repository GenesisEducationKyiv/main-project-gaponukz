package providers

import (
	"btcapp/src/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func fromSymbolToWord(symbol entities.Symbol) (string, error) {
	currencies := map[string]string{
		"BTC": "bitcoin",
		"ETH": "ethereum",
	}

	word, ok := currencies[string(symbol)]
	if !ok {
		return "", fmt.Errorf("symbol %s not supported", symbol)
	}

	return word, nil
}

func loadJsonResponseBody(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
