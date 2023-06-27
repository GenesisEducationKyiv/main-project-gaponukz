package unit

import (
	"btcapp/src/entities"
	"btcapp/src/storage"
	"os"
	"testing"
)

type db interface {
	GetAll() ([]entities.User, error)
	Create(entities.User) error
}

var testUsers []entities.User = []entities.User{
	{Gmail: "Alice"},
	{Gmail: "Bob"},
	{Gmail: "Carol"},
}

func TestJsonDatabase(t *testing.T) {
	testFilename := "test.json"
	err := os.WriteFile(testFilename, []byte("[]"), 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		os.Remove(testFilename)
	}()

	s := storage.NewJsonFileUserStorage(testFilename)

	checkEmptiness(s, t)
	checkCreations(s, t)
	checkExisting(s, t)
}

func checkEmptiness(s db, t *testing.T) {
	users, err := s.GetAll()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(users))
	}
}

func checkCreations(s db, t *testing.T) {
	for _, user := range testUsers {
		err := s.Create(user)
		if err != nil {
			t.Errorf("Unexpected error creating user %s: %v", user.Gmail, err)
		}
	}
}

func checkExisting(s db, t *testing.T) {
	users, err := s.GetAll()
	if err != nil {
		t.Errorf("Unexpected error stogate.GetAll: %v", err)
	}

	for _, testUser := range testUsers {
		var u *entities.User

		for _, user := range users {
			if testUser.Gmail == user.Gmail {
				u = &testUser
				break
			}
		}

		if u == nil {
			t.Errorf("could not find user %s in database", testUser.Gmail)
		}
	}
}
