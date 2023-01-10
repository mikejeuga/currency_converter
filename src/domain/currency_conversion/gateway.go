package currency_conversion

import (
	"github.com/mikejeuga/currency_converter/models"
	"strconv"
)

//go:generate moq -out mocks/conversioner_moq.go -pkg=mocks . Conversioner
type Conversioner interface {
	GetFXRate(base, foreign string) (models.Rate, error)
}

type Gateway struct {
	client  Conversioner
	service *Service
}

func NewGateway(client Conversioner, service *Service) *Gateway {
	return &Gateway{client: client, service: service}
}

func (g *Gateway) GetFXRate(base, foreign string) (models.Rate, error) {
	return g.client.GetFXRate(base, foreign)
}

func (g *Gateway) Convert(amount, baseCurrency, foreignCurrency string) (models.Amount, error) {
	rate, err := g.client.GetFXRate(baseCurrency, foreignCurrency)
	if err != nil {
		return models.Amount{}, err
	}

	baseAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return models.Amount{}, err
	}
	m := models.Amount{
		Unit: baseAmount,
		Currency: models.Currency{
			Code: baseCurrency,
		},
	}

	return g.service.Convert(m, foreignCurrency, rate), nil
}
