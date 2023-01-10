//+go:build unit

package currency_conversion_test

import (
	"github.com/alecthomas/assert"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion/mocks"
	"testing"
)

func TestGateway(t *testing.T) {
	deps := CreateDeps()
	gateway := currency_conversion.NewGateway(deps.ConversionerMock, currency_conversion.NewService())

	expectedFXRate := 0.92
	givenGetRateWasCalled(deps, expectedFXRate)

	rate, err := gateway.GetRate(models.GBP, models.USD)
	assert.NoError(t, err)
	assert.Equal(t, expectedFXRate, rate.Spot)
}

func givenGetRateWasCalled(deps Deps, fxRate float64) {
	deps.ConversionerMock.GetFXRateFunc = func(base string, foreign string) (models.Rate, error) {
		return models.Rate{
			Spot: fxRate,
		}, nil
	}
}

type Deps struct {
	ConversionerMock *mocks.ConversionerMock
}

func CreateDeps() Deps {
	return Deps{
		ConversionerMock: &mocks.ConversionerMock{},
	}
}
