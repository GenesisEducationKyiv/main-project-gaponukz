package mocks

import "btcapp/src/entities"

type exporterStub struct {
	expected float64
}

func (m exporterStub) CurrentRate(from entities.Symbol, to entities.Symbol) (entities.Price, error) {
	return entities.Price(m.expected), nil
}

func NewExporterStub(expected float64) exporterStub {
	return exporterStub{expected: expected}
}
