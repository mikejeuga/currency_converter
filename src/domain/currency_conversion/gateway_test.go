//+go:build unit

package currency_conversion_test

import (
	"github.com/alecthomas/assert"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	mocks2 "github.com/mikejeuga/currency_converter/src/web/mocks"
	"testing"
)

func TestGateway(t *testing.T) {
	deps := CreateDeps()
	gateway := currency_conversion.NewGateway(deps.GatewayMock, currency_conversion.NewService())

	expectedFXRate := 0.92
	givenGetRateWasCalled(deps, expectedFXRate)

	rate, err := gateway.GetRate(models.GBP, models.USD)
	assert.NoError(t, err)
	assert.Equal(t, expectedFXRate, rate.Spot)
}

func givenGetRateWasCalled(deps Deps, fxRate float64) {
	deps.GatewayMock.GetFXRateFunc = func(base string, foreign string) (models.Rate, error) {
		return models.Rate{
			Spot: fxRate,
		}, nil
	}
}

type Deps struct {
	GatewayMock *mocks2.GatewayMock
}

func CreateDeps() Deps {
	return Deps{
		GatewayMock: &mocks2.GatewayMock{},
	}
}
