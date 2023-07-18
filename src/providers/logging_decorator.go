package providers

import (
	"btcapp/src/entities"
	"fmt"
)

type logger interface {
	Info(string)
	Error(string)
}

type loggingDecorator struct {
	provider baseProvider
	logger   logger
}

func NewLoggingDecorator(p baseProvider, l logger) loggingDecorator {
	return loggingDecorator{provider: p, logger: l}
}

func (d loggingDecorator) CurrentRate(from entities.Symbol, to entities.Symbol) (entities.Price, error) {
	rate, err := d.provider.CurrentRate(from, to)
	name := d.provider.Name()

	if err != nil {
		d.logger.Error(fmt.Sprintf("could not get rate with %s because of %v", name, err))
	} else {
		d.logger.Info(fmt.Sprintf("current BTC price according to %s is %f", name, rate))
	}

	return rate, err
}

func (d loggingDecorator) Name() string {
	return d.provider.Name()
}

func (d loggingDecorator) SetNext(next baseProvider) {
	d.provider.SetNext(next)
}
