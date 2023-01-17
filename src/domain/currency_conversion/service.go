package currency_conversion

import (
	"github.com/mikejeuga/currency_converter/models"
)

//go:generate moq -out mocks/conversioner_moq.go -pkg=mocks . Conversioner
type Conversioner interface {
	GetFXRate(base, foreign string) (models.Rate, error)
}

type Service struct {
	client Conversioner
}

func NewService(client Conversioner) *Service {
	return &Service{client: client}
}

func (s *Service) GetRate(base, foreign string) (models.Rate, error) {
	rate, err := s.client.GetFXRate(base, foreign)
	if err != nil {
		return models.Rate{}, err
	}
	return rate, nil
}

func (s *Service) Convert(amount float64, baseCurrency, foreignCurrency string) (models.Amount, error) {
	rate, err := s.client.GetFXRate(baseCurrency, foreignCurrency)
	if err != nil {
		return models.Amount{}, err
	}

	m := models.Amount{
		Unit: amount,
		Currency: models.Currency{
			Code: baseCurrency,
		},
	}

	return s.convert(m, foreignCurrency, rate), nil
}

func (s *Service) convert(amount models.Amount, foreignCurrency string, rate models.Rate) models.Amount {
	convertedAmount := amount.Unit * rate.Spot
	return models.Amount{
		Unit: convertedAmount,
		Currency: models.Currency{
			Code: foreignCurrency,
		},
	}

}
