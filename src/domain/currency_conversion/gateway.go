package currency_conversion

import (
	"github.com/mikejeuga/currency_converter/models"
)

//go:generate moq -out mocks/converter_moq.go -pkg=mocks . Converter
type Converter interface {
	GetRate(base, foreign string) (models.Rate, error)
	Convert(amount models.Amount, foreignCurrency string) (models.Amount, error)
}

type Gateway struct {
	client  Converter
	service *Service
}

func NewGateway(client Converter, service *Service) *Gateway {
	return &Gateway{client: client, service: service}
}

func (g *Gateway) GetRate(base, foreign string) (models.Rate, error) {
	return g.client.GetRate(base, foreign)
}

func (g *Gateway) Convert(amount models.Amount, foreignCurrency string) (models.Amount, error) {
	return models.Amount{}, nil
}
