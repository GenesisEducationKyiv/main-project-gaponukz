package integration

import (
	"btcapp/src/entities"
	"btcapp/src/storage"
	"btcapp/src/usecase"
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

	var currentMessage mocks.Message // last "sended" message
	db := storage.NewJsonFileUserStorage(testFilename)
	ex := mocks.NewMockExporter(expectedPrice)
	n := mocks.NewMockNotifier(func(m mocks.Message) { currentMessage = m })
	s := usecase.NewService(db, ex, n)
	testUser1 := entities.User{Gmail: "testuser1"}

	err = s.SubscribeUser(testUser1)
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
	err = s.NotifySubscribers()
	if err != nil {
		t.Error(err.Error())
	}

	if currentMessage.To != user.Gmail {
		t.Errorf("Expected send to %s, got %s", user.Gmail, currentMessage.To)
	}

	expectedBody := fmt.Sprintf("%f", expectedPrice)
	if currentMessage.Body != expectedBody {
		t.Errorf("Expected msg body %s, got %s", expectedBody, currentMessage.Body)
	}
}
