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
	settings := settings.NewDotEnvSettings().Load()
	storage := storage.NewJsonFileUserStorage("users.json")
	priceExporter := exporter.NewCoingeckoExporter()
	notifier := notifier.NewGmailNotifier(settings.Gmail, settings.GmailPassword)
	service := usecase.NewService(storage, priceExporter, notifier)

	contr := controller.NewController(service)
	app := controller.GetApp(contr, settings.Port)

	fmt.Printf("⚡️[server]: Server is running at http://localhost:%s", settings.Port)
	app.ListenAndServe()
}
