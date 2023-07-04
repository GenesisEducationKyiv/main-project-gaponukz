package unit

import (
	"btcapp/src/exporter"
	"testing"
)

type exp interface {
	CurrentBTCPrice() (float64, error)
	Name() string
}

func TestExporters(t *testing.T) {
	exporters := []exp{
		exporter.NewCoingeckoExporter(),
		exporter.NewCoinstatsExporter(),
		exporter.NewKucoinExporter(),
	}

	for _, e := range exporters {
		price, err := e.CurrentBTCPrice()

		if err != nil {
			t.Errorf("%s.CurrentBTCPrice error: %s", e.Name(), err.Error())
		}

		if price <= 0 {
			t.Errorf("%s gives wrong data (price: %f)", e.Name(), price)
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
