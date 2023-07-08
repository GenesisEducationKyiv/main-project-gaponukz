package exporter

import (
	"fmt"
)

type logger interface {
	Info(string)
	Warn(string)
}

type loggingDecorator struct {
	provider baseProvider
	logger   logger
}

func NewLoggingDecorator(p baseProvider, l logger) loggingDecorator {
	return loggingDecorator{provider: p, logger: l}
}

func (d loggingDecorator) CurrentBTCPrice() (float64, error) {
	rate, err := d.provider.CurrentBTCPrice()
	name := d.provider.Name()

	if err != nil {
		d.logger.Warn(fmt.Sprintf("could not get rate with %s because of %v", name, err))
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
