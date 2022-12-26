package currency_conversion_test

import (
	"github.com/alecthomas/assert"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"testing"
)

func TestService(t *testing.T) {
	//GIVEN a base amount and an exchange rate
	amount := models.Amount{
		Unit: 1000,
		Currency: models.Currency{
			Code: models.GBP,
		},
	}

	exchangeRate := models.Rate{
		Spot: 0.92,
	}

	service := currency_conversion.NewService()

	convertedAmount := service.Convert(amount, exchangeRate)

	expected := convertedAmount.Unit / amount.Unit
	assert.Equal(t, exchangeRate.Spot, expected)
}
