package exporter

import (
	"encoding/json"
	"fmt"
	"io"
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

	body, err := readResponseBody(response.Body)
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

func readResponseBody(body io.Reader) ([]byte, error) {
	var buf []byte
	for {
		tmp := make([]byte, 4096)
		n, err := body.Read(tmp)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	return buf, nil
}
