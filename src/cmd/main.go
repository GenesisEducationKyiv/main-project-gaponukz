package main

import (
	"fmt"

	"btcapp/src/controller"
	"btcapp/src/exporter"
	"btcapp/src/notifier"
	"btcapp/src/settings"
	"btcapp/src/storage"
	"btcapp/src/usecase"
)

func main() {
	baseRateProvider := exporter.NewCoingeckoExporter()
	coinstatsProviderHelper := exporter.NewCoinstatsExporter()
	kukoinProviderHelper := exporter.NewKucoinExporter()
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
