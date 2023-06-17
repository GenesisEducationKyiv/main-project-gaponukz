package main

import (
	"fmt"
	"net/http"

	. "btcapp/src/controller"
	"btcapp/src/exporter"
	"btcapp/src/observer"
	"btcapp/src/settings"
	"btcapp/src/storage"
)

func main() {
	settingsExporter := settings.DotEnvSettings{}
	settings := settingsExporter.Load()

	routerService := Controller{
		Exporter: exporter.CoingeckoExporter{},
		Storage: storage.JsonFileUserStorage{
			Filename: "users.json",
		},
		Observer: observer.ObserverService{
			Strategy: func(userGmail string, BTCPrice float64) error {
				return observer.SendGmailLetterStrategy(
					userGmail, BTCPrice, settings,
				)
			},
		},
	}

	httpRoute := http.NewServeMux()

	httpRoute.HandleFunc("/rate", RequiredMethod(routerService.RateRouter, http.MethodGet))
	httpRoute.HandleFunc("/subscribe", RequiredMethod(routerService.SubscribeRouter, http.MethodPost))
	httpRoute.HandleFunc("/sendEmails", RequiredMethod(routerService.SendEmailsRouter, http.MethodPost))

	loggedRouter := LoggingMiddleware(httpRoute)

	server := http.Server{
		Addr:    ":8080",
		Handler: loggedRouter,
	}

	fmt.Println("⚡️[server]: Server is running at http://localhost:8080")
	server.ListenAndServe()
}
