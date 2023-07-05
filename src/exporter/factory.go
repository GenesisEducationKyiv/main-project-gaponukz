package exporter

type fetchRateFunction func() (float64, error)

type baseProvider interface {
	CurrentBTCPrice() (float64, error)
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

func (e *provider) CurrentBTCPrice() (float64, error) {
	rate, err := e.fetchRate()
	if err == nil {
		return rate, nil
	}

	if e.next == nil {
		return 0, err
	}

	return e.next.CurrentBTCPrice()
}

func (e *provider) Name() string {
	return e.name
}

func (e *provider) SetNext(next baseProvider) {
	e.next = next
}
