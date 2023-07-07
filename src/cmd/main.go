package main

import (
	"fmt"

	"btcapp/src/controller"
	"btcapp/src/logger"
	gmailNotifier "btcapp/src/notifier"
	"btcapp/src/providers"
	"btcapp/src/settings"
	"btcapp/src/storage"
	"btcapp/src/usecases/currency_rate"
	"btcapp/src/usecases/notifier"
	"btcapp/src/usecases/subscription"
)

func main() {
	logger := logger.NewConsoleLogger()
	baseRateProvider := providers.NewCoingeckoProvider()
	coinstatsProviderHelper := providers.NewCoinstatsProvider()
	kukoinProviderHelper := providers.NewKucoinProvider()

	baseRateProvider = providers.NewLoggingDecorator(baseRateProvider, logger)
	coinstatsProviderHelper = providers.NewLoggingDecorator(coinstatsProviderHelper, logger)
	kukoinProviderHelper = providers.NewLoggingDecorator(kukoinProviderHelper, logger)

	baseRateProvider.SetNext(coinstatsProviderHelper)
	coinstatsProviderHelper.SetNext(kukoinProviderHelper)

	settings := settings.NewDotEnvSettings().Load()
	storage := storage.NewJsonFileUserStorage("users.json")
	gn := gmailNotifier.NewGmailNotifier(settings.GmailServer, settings.Gmail, settings.GmailPassword)
	loggeredNotifier := gmailNotifier.NewLoggingDecorator(gn, logger)

	rateService := currency_rate.NewCurrencyRateService(baseRateProvider)
	notifierService := notifier.NewNotifierService(loggeredNotifier)
	subscriptionService := subscription.NewSubscriptionService(storage)

	contr := controller.NewController(rateService, notifierService, subscriptionService)
	app := controller.GetApp(contr, settings.Port)

	fmt.Printf("⚡️[server]: Server is running at http://localhost:%s", settings.Port)

	err := app.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
