package controller

import (
	"net/http"
	"strconv"

	"btcapp/src/entities"
)

type service interface {
	GetCurrentPrice() (float64, error)
	SubscribeUser(entities.User) error
	NotifySubscribers() error
}

type controller struct {
	service service
}

func NewController(service service) *controller {
	return &controller{service: service}
}

func (c controller) RateRouter(responseWriter http.ResponseWriter, request *http.Request) {
	btcPrice, err := c.service.GetCurrentPrice()

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	stringPrice := strconv.FormatFloat(btcPrice, 'f', -1, 64)
	responseWriter.Write([]byte(stringPrice))
}

func (c controller) SubscribeRouter(responseWriter http.ResponseWriter, request *http.Request) {
	userGmail := request.URL.Query().Get("gmail")
	if userGmail == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	err := c.service.SubscribeUser(entities.User{
		Gmail: userGmail,
	})

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Write([]byte("Added"))
}

func (c controller) SendEmailsRouter(responseWriter http.ResponseWriter, request *http.Request) {
	err := c.service.NotifySubscribers()
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Write([]byte("Sended"))
}
