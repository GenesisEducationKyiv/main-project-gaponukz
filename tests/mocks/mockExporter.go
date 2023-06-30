package mocks

type exporterStub struct {
	expected float64
}

func (m exporterStub) CurrentBTCPrice() (float64, error) {
	return m.expected, nil
}

func NewExporterStub(expected float64) exporterStub {
	return exporterStub{expected: expected}
}
