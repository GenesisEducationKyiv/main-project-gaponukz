package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type coingeckoExporter struct{}

func NewCoingeckoExporter() *coingeckoExporter {
	return &coingeckoExporter{}
}

func (e coingeckoExporter) GetCurrentBTCPrice() (float64, error) {
	var apiResponse map[string]map[string]float64
	const ApiUrl = "https://api.coingecko.com/api/v3/simple/price"
	response, err := http.Get(fmt.Sprintf("%s?ids=bitcoin&vs_currencies=uah", ApiUrl))

	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return 0, err
	}

	price, ok := apiResponse["bitcoin"]["uah"]

	if !ok {
		return 0, fmt.Errorf("api response error")
	}

	return price, nil
}
