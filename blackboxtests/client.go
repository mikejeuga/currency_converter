package blackboxtests

import (
	"github.com/mikejeuga/currency_converter/models"
	"net/http"
	"time"
)

type TestUser struct {
	baseURL string
	client  *http.Client
}

func NewTestUser(baseUrl string) *TestUser {
	c := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}
	return &TestUser{baseURL: baseUrl, client: c}
}

func (t *TestUser) GetRate(base, foreign string) (models.Rate, error) {
	rate := models.Rate{
		Spot: 0,
	}
	return rate, nil
}

func (t *TestUser) Convert(amount models.Amount, rate models.Rate) (models.Amount, error) {
	return models.Amount{
		Unit: 0,
		Currency: models.Currency{
			Code: models.GBP,
		},
	}, nil
}
