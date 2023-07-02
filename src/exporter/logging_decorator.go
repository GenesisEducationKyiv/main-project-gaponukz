package exporter

import (
	"fmt"
)

type logger interface {
	Info(string)
	Warn(string)
}

type baseProvider interface {
	CurrentBTCPrice() (float64, error)
}

type loggingDecorator struct {
	providerName string
	provider     baseProvider
	logger       logger
}

func NewLoggingDecorator(n string, p baseProvider, l logger) loggingDecorator {
	return loggingDecorator{providerName: n, provider: p, logger: l}
}

func (d loggingDecorator) CurrentBTCPrice() (float64, error) {
	rate, err := d.provider.CurrentBTCPrice()
	if err != nil {
		d.logger.Warn(fmt.Sprintf("could not get rate with %s because of %v", d.providerName, err))
	} else {
		d.logger.Info(fmt.Sprintf("current BTC price according to %s is %f", d.provider, rate))
	}

	return rate, err
}
