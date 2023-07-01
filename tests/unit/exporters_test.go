package unit

import (
	"btcapp/src/exporter"
	"testing"
)

type exp interface {
	CurrentBTCPrice() (float64, error)
}

func TestExporters(t *testing.T) {
	exporters := []struct {
		name     string
		provider exp
	}{
		{name: "Coingecko", provider: exporter.NewCoingeckoExporter()},
		{name: "Coinstats", provider: exporter.NewCoinstatsExporter()},
		{name: "Kucoin", provider: exporter.NewKucoinExporter()},
	}

	for _, e := range exporters {
		price, err := e.provider.CurrentBTCPrice()

		if err != nil {
			t.Errorf("%s.CurrentBTCPrice error: %s", e.name, err.Error())
		}

		if price <= 0 {
			t.Errorf("%s gives wrong data (price: %f)", e.name, price)
		}
	}
}

func TestChainOfExporters(t *testing.T) {
	baseRateProvider := exporter.NewCoingeckoExporter()
	coinstatsProviderHelper := exporter.NewCoinstatsExporter()
	kukoinProviderHelper := exporter.NewKucoinExporter()
	baseRateProvider.SetNext(coinstatsProviderHelper)
	coinstatsProviderHelper.SetNext(kukoinProviderHelper)

	price, err := baseRateProvider.CurrentBTCPrice()

	if err != nil {
		t.Error(err.Error())
	}

	if price <= 0 {
		t.Errorf("gives wrong data (price: %f)", price)
	}
}
