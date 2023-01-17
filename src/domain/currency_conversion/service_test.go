//+go:build unit

package currency_conversion_test

import (
	"github.com/alecthomas/assert/v2"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"testing"
)

func TestGateway(t *testing.T) {
	client := NewClientStub(0.92)
	gateway := currency_conversion.NewService(client)

	rate, err := gateway.GetRate(models.GBP, models.USD)
	assert.NoError(t, err)
	assert.Equal(t, client.Spot.Spot, rate.Spot)

}

type ClientStub struct {
	Spot models.Rate
}

func NewClientStub(spot float64) *ClientStub {

	return &ClientStub{Spot: models.Rate{
		Spot: spot,
	}}
}

func (c ClientStub) GetFXRate(base, foreign string) (models.Rate, error) {
	return c.Spot, nil
}
