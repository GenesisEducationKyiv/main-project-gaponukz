package controller

import (
	"fmt"
	"log"
	"net/http"
)

func GetApp(c *controller, port string) *http.Server {
	httpRoute := http.NewServeMux()

	httpRoute.HandleFunc("/rate", requiredMethod(c.RateRouter, http.MethodGet))
	httpRoute.HandleFunc("/subscribe", requiredMethod(c.SubscribeRouter, http.MethodPost))
	httpRoute.HandleFunc("/sendEmails", requiredMethod(c.SendEmailsRouter, http.MethodPost))

	loggedRouter := loggingMiddleware(httpRoute)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: loggedRouter,
	}
}

type routerFunc = func(rw http.ResponseWriter, r *http.Request)

func requiredMethod(router routerFunc, required string) routerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == required {
			router(responseWriter, request)

		} else {
			http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s?%s", request.Method, request.URL.Path, request.URL.RawQuery)
		next.ServeHTTP(responseWriter, request)
	})
}
