//+go:build unit

package currency_conversion_test

import (
	"github.com/alecthomas/assert"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"testing"
)

func TestGateway(t *testing.T) {
	client := &ClientStub{}
	gateway := currency_conversion.NewService(client)

	expectedFXRate, err := client.GetFXRate(models.GBP, models.USD)
	assert.NoError(t, err)

	rate, err := gateway.GetRate(models.GBP, models.USD)
	assert.NoError(t, err)
	assert.Equal(t, expectedFXRate.Spot, rate.Spot)

}

type ClientStub struct{}

func (c ClientStub) GetFXRate(base, foreign string) (models.Rate, error) {
	return models.Rate{
		Spot: 0.92,
	}, nil
}
