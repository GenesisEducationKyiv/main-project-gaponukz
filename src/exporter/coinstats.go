package exporter

import "fmt"

type coinstatsExporter struct{}

func NewCoinstatsExporter() coinstatsExporter {
	return coinstatsExporter{}
}

type coinstatsResponse struct {
	Coin struct {
		ID              string   `json:"id"`
		Icon            string   `json:"icon"`
		Name            string   `json:"name"`
		Symbol          string   `json:"symbol"`
		Rank            int      `json:"rank"`
		Price           float64  `json:"price"`
		PriceBtc        float64  `json:"priceBtc"`
		Volume          float64  `json:"volume"`
		MarketCap       float64  `json:"marketCap"`
		AvailableSupply float64  `json:"availableSupply"`
		TotalSupply     float64  `json:"totalSupply"`
		PriceChange1h   float64  `json:"priceChange1h"`
		PriceChange1d   float64  `json:"priceChange1d"`
		PriceChange1w   float64  `json:"priceChange1w"`
		WebsiteUrl      string   `json:"websiteUrl"`
		TwitterUrl      string   `json:"twitterUrl"`
		Exp             []string `json:"exp"`
	} `json:"coin"`
}

func (e coinstatsExporter) CurrentBTCPrice() (float64, error) {
	var apiResponse coinstatsResponse
	const ApiUrl = "https://api.coinstats.app/public/v1/coins/bitcoin"

	err := getJson(fmt.Sprintf("%s?currency=UAH", ApiUrl), &apiResponse)
	if err != nil {
		return 0, err
	}

	return apiResponse.Coin.Price, nil
}
