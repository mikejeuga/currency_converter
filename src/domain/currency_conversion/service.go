package currency_conversion

import (
	"github.com/mikejeuga/currency_converter/models"
)

type Service struct{}

func NewService() *Service {
	return &Service{}

}

func (s *Service) Convert(amount models.Amount, rate models.Rate) models.Amount {
	convertedAmount := amount.Unit * rate.Spot
	return models.Amount{
		Unit: convertedAmount,
		Currency: models.Currency{
			Code: amount.Currency.Code,
		},
	}
}
