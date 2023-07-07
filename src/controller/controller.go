package controller

import (
	"net/http"
	"strconv"

	"btcapp/src/entities"
)

type rateService interface {
	CurrentBTCRate() (entities.Price, error)
}

type notifierService interface {
	NotifyBTCPrice(users []entities.User, price entities.Price)
}

type subscriptionService interface {
	SubscribeUser(user entities.User) error
	All(filter func(entities.User) bool) ([]entities.User, error)
}

type controller struct {
	rs rateService
	ns notifierService
	ss subscriptionService
}

func NewController(rs rateService, ns notifierService, ss subscriptionService) *controller {
	return &controller{rs: rs, ns: ns, ss: ss}
}

func (c controller) RateRouter(responseWriter http.ResponseWriter, request *http.Request) {
	btcPrice, err := c.rs.CurrentBTCRate()

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	stringPrice := strconv.FormatFloat(float64(btcPrice), 'f', -1, 64)
	responseWriter.Write([]byte(stringPrice))
}

func (c controller) SubscribeRouter(responseWriter http.ResponseWriter, request *http.Request) {
	userGmail := request.URL.Query().Get("gmail")
	if userGmail == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	err := c.ss.SubscribeUser(entities.User{
		Gmail: userGmail,
	})

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Write([]byte("Added"))
}

func (c controller) SendEmailsRouter(responseWriter http.ResponseWriter, request *http.Request) {
	users, err := c.ss.All(nil)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	price, err := c.rs.CurrentBTCRate()
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.ns.NotifyBTCPrice(users, price)
	responseWriter.Write([]byte("Sended"))
}
