package observer

import (
	"btcapp/src/entities"
	exporterService "btcapp/src/exporter"
)

type GetAllAbleStorage interface {
	GetAll() ([]entities.User, error)
}

type IObserverService interface {
	Notify(
		exporter exporterService.IRateExporter,
		storage GetAllAbleStorage,
	) error
}

type ObserverService struct {
	Strategy func(userGmail string, BTCPrice float64) error
}

func (obs ObserverService) Notify(
	exporter exporterService.IRateExporter,
	storage GetAllAbleStorage,
) error {
	users, err := storage.GetAll()
	if err != nil {
		return err
	}

	btcPrice, err := exporter.GetCurrentBTCPrice()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.HasSubscription {
			obs.Strategy(user.Gmail, btcPrice)
		}
	}

	return nil
}
