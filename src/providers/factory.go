package providers

import "btcapp/src/entities"

type fetchRateFunction func(entities.Symbol, entities.Symbol) (entities.Price, error)

type baseProvider interface {
	CurrentRate(entities.Symbol, entities.Symbol) (entities.Price, error)
	SetNext(baseProvider)
	Name() string
}

type provider struct {
	name      string
	next      baseProvider
	fetchRate fetchRateFunction
}

func providerChainFactory(name string, fetchRate fetchRateFunction) baseProvider {
	return &provider{name: name, fetchRate: fetchRate}
}

func (e *provider) CurrentRate(from entities.Symbol, to entities.Symbol) (entities.Price, error) {
	rate, err := e.fetchRate(from, to)
	if err == nil {
		return rate, nil
	}

	if e.next == nil {
		return 0, err
	}

	return e.next.CurrentRate(from, to)
}

func (e *provider) Name() string {
	return e.name
}

func (e *provider) SetNext(next baseProvider) {
	e.next = next
}
