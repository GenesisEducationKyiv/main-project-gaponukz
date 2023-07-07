package currency_rate

import "btcapp/src/entities"

type provider interface {
	CurrentRate(entities.Symbol, entities.Symbol) (entities.Price, error)
}

type currencyRateService struct {
	p provider
}

func NewCurrencyRateService(p provider) currencyRateService {
	return currencyRateService{p: p}
}

func (s currencyRateService) CurrentBTCRate() (entities.Price, error) {
	return s.p.CurrentRate("BTC", "UAH")
}
