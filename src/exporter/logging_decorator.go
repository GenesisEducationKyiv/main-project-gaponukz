package exporter

import (
	"fmt"
)

type logger interface {
	Info(string)
	Warn(string)
}

type rateProvider interface {
	CurrentBTCPrice() (float64, error)
	Name() string
}

type loggingDecorator struct {
	provider rateProvider
	logger   logger
}

func NewLoggingDecorator(n string, p rateProvider, l logger) loggingDecorator {
	return loggingDecorator{provider: p, logger: l}
}

func (d loggingDecorator) CurrentBTCPrice() (float64, error) {
	rate, err := d.provider.CurrentBTCPrice()
	name := d.provider.Name()

	if err != nil {
		d.logger.Warn(fmt.Sprintf("could not get rate with %s because of %v", name, err))
	} else {
		d.logger.Info(fmt.Sprintf("current BTC price according to %s is %f", d.provider, rate))
	}

	return rate, err
}
