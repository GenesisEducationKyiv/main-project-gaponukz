package integration

import (
	"btcapp/src/entities"
	"btcapp/src/file_storage"
	"btcapp/src/usecases/currency_rate"
	"btcapp/src/usecases/notifier"
	"btcapp/src/usecases/subscription"
	"btcapp/tests/mocks"
	"fmt"
	"os"
	"testing"

	"golang.org/x/exp/slices"
)

func TestIntegration(t *testing.T) {
	const testFilename = "test.json"
	const expectedPrice = 69.69

	err := os.WriteFile(testFilename, []byte("[]"), 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		os.Remove(testFilename)
	}()

	var lastSendedMessage mocks.Message
	db := file_storage.NewJsonFileUserStorage(testFilename)
	ex := mocks.NewExporterStub(expectedPrice)
	n := mocks.NewMockNotifier(func(m mocks.Message) { lastSendedMessage = m })

	rateService := currency_rate.NewCurrencyRateService(ex)
	notifierService := notifier.NewNotifierService(n)
	subscriptionService := subscription.NewSubscriptionService(db)

	testUser1 := entities.User{Gmail: "testuser1"}

	price, err := rateService.CurrentBTCRate()
	if err != nil {
		t.Error(err.Error())
	}

	err = subscriptionService.SubscribeUser(testUser1)
	if err != nil {
		t.Error(err.Error())
	}

	users, err := db.GetAll()
	if err != nil {
		t.Error(err.Error())
	}

	index := slices.IndexFunc(users, func(u entities.User) bool { return u.Gmail == testUser1.Gmail })
	if index == -1 {
		t.Error("user not found in database after subscription")
	}

	user := users[index]

	notifierService.NotifyBTCPrice(users, entities.Price(price))
	if err != nil {
		t.Error(err.Error())
	}

	if lastSendedMessage.To != user.Gmail {
		t.Errorf("Expected send to %s, got %s", user.Gmail, lastSendedMessage.To)
	}

	expectedBody := fmt.Sprintf("%f", expectedPrice)
	if lastSendedMessage.Body != expectedBody {
		t.Errorf("Expected msg body %s, got %s", expectedBody, lastSendedMessage.Body)
	}
}
