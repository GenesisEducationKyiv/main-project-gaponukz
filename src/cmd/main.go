package main

import (
	"fmt"

	"btcapp/src/controller"
	"btcapp/src/exporter"
	"btcapp/src/logger"
	"btcapp/src/notifier"
	"btcapp/src/settings"
	"btcapp/src/storage"
)

func main() {
	logger := logger.NewConsoleLogger()
	baseRateProvider := exporter.NewCoingeckoExporter()
	coinstatsProviderHelper := exporter.NewCoinstatsExporter()
	kukoinProviderHelper := exporter.NewKucoinExporter()

	baseRateProvider = exporter.NewLoggingDecorator(baseRateProvider, logger)
	coinstatsProviderHelper = exporter.NewLoggingDecorator(coinstatsProviderHelper, logger)
	kukoinProviderHelper = exporter.NewLoggingDecorator(kukoinProviderHelper, logger)

	baseRateProvider.SetNext(coinstatsProviderHelper)
	coinstatsProviderHelper.SetNext(kukoinProviderHelper)

	settings := settings.NewDotEnvSettings().Load()
	storage := storage.NewJsonFileUserStorage("users.json")
	notifier := notifier.NewGmailNotifier(settings.GmailServer, settings.Gmail, settings.GmailPassword)

	service := usecase.NewService(storage, baseRateProvider, notifier)

	contr := controller.NewController(service)
	app := controller.GetApp(contr, settings.Port)

	fmt.Printf("⚡️[server]: Server is running at http://localhost:%s", settings.Port)

	err := app.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
