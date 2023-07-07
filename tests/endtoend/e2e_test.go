package endtoend

import (
	"btcapp/src/controller"
	"btcapp/src/storage"
	"btcapp/src/usecases/currency_rate"
	"btcapp/src/usecases/notifier"
	"btcapp/src/usecases/subscription"
	"btcapp/tests/mocks"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

const expectedPrice = 69.69

func TestHTTPRoutes(t *testing.T) {
	const freePort = 3000
	const testFilename = "test.json"
	err := os.WriteFile(testFilename, []byte("[]"), 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		os.Remove(testFilename)
	}()

	go startTestServer(freePort, testFilename)

	testCases := []struct {
		method         string
		endpoint       string
		expectedResult string
	}{
		{"get", "rate", strconv.FormatFloat(expectedPrice, 'f', -1, 64)},
		{"post", "subscribe?gmail=testuser", "Added"},
		{"post", "sendEmails", "Sended"},
	}

	for _, tc := range testCases {
		var err error
		var response string

		if tc.method == "post" {
			response, err = postBody(formatUrl(tc.endpoint, freePort))

		} else if tc.method == "get" {
			response, err = getBody(formatUrl(tc.endpoint, freePort))
		}

		if err != nil {
			t.Fatal(err.Error())
		}

		if response != tc.expectedResult {
			t.Errorf("for endpoint '%s', expected: %s, got: %s", tc.endpoint, tc.expectedResult, response)
		}
	}
}

func startTestServer(freePort int, testFilename string) {
	db := storage.NewJsonFileUserStorage(testFilename)
	ex := mocks.NewExporterStub(expectedPrice)
	n := mocks.NewMockNotifier(func(m mocks.Message) {})

	rateService := currency_rate.NewCurrencyRateService(ex)
	notifierService := notifier.NewNotifierService(n)
	subscriptionService := subscription.NewSubscriptionService(db)

	contr := controller.NewController(rateService, notifierService, subscriptionService)
	app := controller.GetApp(contr, fmt.Sprintf("%d", freePort))

	err := app.ListenAndServe()
	if err != nil {
		panic(err.Error())
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
