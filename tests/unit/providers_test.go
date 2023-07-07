package unit

import (
	"btcapp/src/entities"
	"btcapp/src/providers"
	"testing"
)

type p interface {
	CurrentRate(entities.Symbol, entities.Symbol) (entities.Price, error)
	Name() string
}

func TestExporters(t *testing.T) {
	providers := []p{
		providers.NewCoingeckoProvider(),
		providers.NewCoinstatsProvider(),
		providers.NewKucoinProvider(),
	}

	for _, e := range providers {
		price, err := e.CurrentRate("BTC", "UAH")

		if err != nil {
			t.Errorf("%s.CurrentRate error: %s", e.Name(), err.Error())
		}

		if price <= 0 {
			t.Errorf("%s gives wrong data (price: %f)", e.Name(), price)
		}
	}
}

func TestChainOfProviders(t *testing.T) {
	baseRateProvider := providers.NewCoingeckoProvider()
	coinstatsProviderHelper := providers.NewCoinstatsProvider()
	kukoinProviderHelper := providers.NewKucoinProvider()
	baseRateProvider.SetNext(coinstatsProviderHelper)
	coinstatsProviderHelper.SetNext(kukoinProviderHelper)

	price, err := baseRateProvider.CurrentRate("BTC", "UAH")

	if err != nil {
		t.Error(err.Error())
	}

	if price <= 0 {
		t.Errorf("gives wrong data (price: %f)", price)
	}
}
