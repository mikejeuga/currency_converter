package specifications

import (
	"github.com/alecthomas/assert"
	"github.com/mikejeuga/currency_converter/models"
	"testing"
)

type Converter interface {
	GetFXRate(base, foreign string) (models.Rate, error)
	Convert(amount models.Amount, foreignCurrency string) (models.Amount, error)
}

type CurrencyConversionSpec struct {
	converter Converter
}

func NewCurrencyConversionSpec(converter Converter) *CurrencyConversionSpec {
	return &CurrencyConversionSpec{converter: converter}
}

func (s *CurrencyConversionSpec) CanConverterBaseToForeign(t *testing.T) {
	//Given a base amount (GBP) in my bank account
	amountToConvert := models.Amount{
		Unit: 2000,
		Currency: models.Currency{
			Code: models.GBP,
		},
	}

	//When I need to convert it in USD
	convertedAmount, err := s.converter.Convert(amountToConvert, models.USD)
	assert.NoError(t, err)

	//Then I get a converted amount at the correct FX rate
	rate, err := s.converter.GetFXRate(models.GBP, models.USD)
	assert.NoError(t, err)

	assert.Equal(t, convertedAmount.Unit/amountToConvert.Unit, rate.Spot)
}
