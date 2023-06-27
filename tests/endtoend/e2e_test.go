package endtoend

import (
	"btcapp/src/controller"
	"btcapp/src/storage"
	"btcapp/src/usecase"
	"btcapp/tests/mocks"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

const expectedPrice = 69.69

func TestEndToEnd(t *testing.T) {
	const freePort = 6969
	const testFilename = "test.json"
	err := os.WriteFile(testFilename, []byte("[]"), 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		os.Remove(testFilename)
	}()

	go startTestServer(freePort, testFilename)

	response, err := getBody(formatUrl("rate", freePort))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStringPrice := strconv.FormatFloat(expectedPrice, 'f', -1, 64)
	if expectedStringPrice != response {
		t.Errorf("expected rate: %s, got %s", expectedStringPrice, response)
	}

	response, err = postBody(formatUrl("subscribe?gmail=testuser", freePort))
	if err != nil {
		t.Fatal(err.Error())
	}

	if response != "Added" {
		t.Errorf("after subscription expect: Added, got: %s", response)
	}

	response, err = postBody(formatUrl("sendEmails", freePort))
	if err != nil {
		t.Fatal(err.Error())
	}

	if response != "Sended" {
		t.Errorf("after subscription expect: Sended, got: %s", response)
	}

	// TODO: check if user get gmail
}

func startTestServer(freePort int, testFilename string) {
	storage := storage.NewJsonFileUserStorage(testFilename)
	priceExporter := mocks.NewMockExporter(expectedPrice)
	notifier := mocks.NewMockNotifier(func(m mocks.Message) {})
	service := usecase.NewService(storage, priceExporter, notifier)

	contr := controller.NewController(service)
	app := controller.GetApp(contr, fmt.Sprintf("%d", freePort))

	err := app.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getBody(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func postBody(url string) (string, error) {
	response, err := http.Post(url, "", nil)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func formatUrl(path string, port int) string {
	return fmt.Sprintf("http://localhost:%d/%s", port, path)
}
