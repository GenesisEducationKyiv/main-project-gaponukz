package unit

type mockExporter struct {
	expected float64
}

func (m mockExporter) GetCurrentBTCPrice() (float64, error) {
	return m.expected, nil
}

func NewMockExporter(expected float64) mockExporter {
	return mockExporter{expected: expected}
}
