package unit

import (
	"btcapp/src/entities"
	"btcapp/src/usecases/currency_rate"
	"btcapp/src/usecases/notifier"
	"btcapp/src/usecases/subscription"
	"btcapp/tests/mocks"
	"fmt"
	"testing"
)

func TestSubscriptionService(t *testing.T) {
	storage := mocks.NewStorageMock()
	service := subscription.NewSubscriptionService(storage)
	testUser := entities.User{Gmail: "test1"}

	err := service.SubscribeUser(testUser)
	if err != nil {
		t.Error(err.Error())
	}

	if !storage.Contains(testUser) {
		t.Error("after subscription, the user is not saved to storage")
	}

	err = service.SubscribeUser(testUser)
	if err == nil {
		t.Error("can subscribe the same user twice")
	}
}

func TestCurrencyRateService(t *testing.T) {
	const expected = 69.69
	provider := mocks.NewExporterStub(expected)
	service := currency_rate.NewCurrencyRateService(provider)

	price, err := service.CurrentBTCRate()

	if err != nil {
		t.Errorf("Errog geting price: %s", err.Error())
	}

	if price != expected {
		t.Errorf("Expected %f, got %f", expected, price)
	}
}

func TestNotifySubscribers(t *testing.T) {
	var mess mocks.Message
	const expected = 69.69

	n := mocks.NewMockNotifier(func(m mocks.Message) { mess = m })
	service := notifier.NewNotifierService(n)
	user := entities.User{Gmail: "test1"}
	users := []entities.User{user}

	service.NotifyBTCPrice(users, entities.Price(expected))

	if mess.To != user.Gmail {
		t.Errorf("Expected send to %s, got %s", user.Gmail, mess.To)
	}

	expectedBody := fmt.Sprintf("%f", expected)
	if mess.Body != expectedBody {
		t.Errorf("Expected msg body %s, got %s", expectedBody, mess.Body)
	}
}
