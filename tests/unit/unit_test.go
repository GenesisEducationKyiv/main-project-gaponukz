package unit

import (
	"btcapp/src/entities"
	"btcapp/src/exporter"
	"btcapp/src/usecase"
	"fmt"
	"testing"
)

func TestCoingeckoExporter(t *testing.T) {
	e := exporter.NewCoingeckoExporter()
	price, err := e.GetCurrentBTCPrice()

	if err != nil {
		t.Errorf("Exporter.GetCurrentBTCPrice error: %s", err.Error())
	}

	if price <= 0 {
		t.Errorf("price exporter gives wrong data (price: %f)", price)
	}
}

func TestSubscribeUser(t *testing.T) {
	storage := NewStorageMock()
	service := usecase.NewService(storage, nil, nil)
	testUser := entities.User{Gmail: "test1"}

	service.SubscribeUser(testUser)

	if !storage.Contains(testUser) {
		t.Error("after subscription, the user is not saved to storage")
	}
}

func TestGetCurrentPrice(t *testing.T) {
	const expected = 69.69
	exporter := NewMockExporter(expected)
	service := usecase.NewService(nil, exporter, nil)

	price, err := service.GetCurrentPrice()

	if err != nil {
		t.Errorf("Errog geting price: %s", err.Error())
	}

	if price != expected {
		t.Errorf("Expected %f, got %f", expected, price)
	}
}

func TestNotifySubscribers(t *testing.T) {
	var mess message
	const expected = 69.69

	user := entities.User{Gmail: "test1"}
	storage := NewStorageMock()
	exporter := NewMockExporter(expected)
	notifier := NewMockNotifier(func(m message) { mess = m })
	service := usecase.NewService(storage, exporter, notifier)

	storage.Create(user)

	err := service.NotifySubscribers()
	if err != nil {
		t.Errorf("Errog notifying price: %s", err.Error())
	}

	if mess.to != user.Gmail {
		t.Errorf("Expected send to %s, got %s", user.Gmail, mess.to)
	}

	expectedBody := fmt.Sprintf("%f", expected)
	if mess.body != expectedBody {
		t.Errorf("Expected msg body %s, got %s", expectedBody, mess.body)
	}
}
