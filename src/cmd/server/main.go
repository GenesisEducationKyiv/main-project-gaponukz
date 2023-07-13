package main

import (
	"fmt"

	"btcapp/src/controller"
	"btcapp/src/file_storage"
	"btcapp/src/gmail_notifier"
	"btcapp/src/providers"
	"btcapp/src/rabbitmq_logger"
	"btcapp/src/settings"
	"btcapp/src/usecases/currency_rate"
	"btcapp/src/usecases/notifier"
	"btcapp/src/usecases/subscription"
)

func main() {
	settings := settings.NewDotEnvSettings().Load()

	logger, err := rabbitmq_logger.NewRabbitMQLogger(settings.RabbitMQUrl)
	if err != nil {
		panic(err.Error())
	}

	baseRateProvider := providers.NewCoingeckoProvider()
	coinstatsProviderHelper := providers.NewCoinstatsProvider()
	kukoinProviderHelper := providers.NewKucoinProvider()

	baseRateProvider = providers.NewLoggingDecorator(baseRateProvider, logger)
	coinstatsProviderHelper = providers.NewLoggingDecorator(coinstatsProviderHelper, logger)
	kukoinProviderHelper = providers.NewLoggingDecorator(kukoinProviderHelper, logger)

	baseRateProvider.SetNext(coinstatsProviderHelper)
	coinstatsProviderHelper.SetNext(kukoinProviderHelper)

	storage := file_storage.NewJsonFileUserStorage("users.json")
	gmailNotifier := gmail_notifier.NewGmailNotifier(settings.GmailServer, settings.Gmail, settings.GmailPassword)
	loggeredNotifier := gmail_notifier.NewLoggingDecorator(gmailNotifier, logger)

	rateService := currency_rate.NewCurrencyRateService(baseRateProvider)
	notifierService := notifier.NewNotifierService(loggeredNotifier)
	subscriptionService := subscription.NewSubscriptionService(storage)

	contr := controller.NewController(rateService, notifierService, subscriptionService)
	app := controller.GetApp(contr, settings.Port)

	fmt.Printf("⚡️[server]: Server is running at http://localhost:%s", settings.Port)

	err = app.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
