//+go:build acceptance

package client

import (
	"encoding/json"
	"fmt"
	"github.com/mikejeuga/currency_converter/config"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/web/auth"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type TestUser struct {
	config config.Config
	client *http.Client
}

func NewTestUser(config config.Config) *TestUser {
	c := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}

	return &TestUser{config: config, client: c}
}

func (u *TestUser) GetFXRate(base, foreign string) (models.Rate, error) {
	rateURL, err := url.JoinPath(u.config.BaseURL, "rate")
	if err != nil {
		return models.Rate{}, err
	}

	req, err := http.NewRequest(http.MethodGet, rateURL, nil)
	if err != nil {
		return models.Rate{}, err
	}

	req.Header.Set(auth.TheApiKey, u.config.ApiKey)

	addQueryParams(req, "1", base, foreign)

	res, err := u.client.Do(req)
	if err != nil {
		return models.Rate{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Rate{}, err
	}

	if string(data) == "" {
		return models.Rate{}, err
	}

	suffix := strings.TrimSuffix(string(data), "\n")

	var rateRes models.Rate
	err = json.Unmarshal([]byte(suffix), &rateRes)
	if err != nil {
		return models.Rate{}, err
	}

	return rateRes, nil
}

func (u *TestUser) Convert(amount models.Amount, foreignCurrency string) (models.Amount, error) {
	conversionURL, err := url.JoinPath(u.config.BaseURL, "converted-amount")
	if err != nil {
		return models.Amount{}, err
	}

	req, err := http.NewRequest(http.MethodGet, conversionURL, nil)
	if err != nil {
		return models.Amount{}, err
	}

	req.Header.Set(auth.TheApiKey, u.config.ApiKey)

	amountStr := fmt.Sprintf("%.2f", amount.Unit)

	addQueryParams(req, amountStr, amount.Currency.Code, foreignCurrency)

	res, err := u.client.Do(req)
	if err != nil {
		return models.Amount{}, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Amount{}, err
	}

	var returnedAmount models.Amount
	err = json.Unmarshal(data, &returnedAmount)
	if err != nil {
		return models.Amount{}, err
	}
	return returnedAmount, nil

}

func addQueryParams(req *http.Request, amount, base, foreign string) {
	q := req.URL.Query()
	if amount != "1" {
		q.Add("amount", amount)
	}
	q.Add("have", base)
	q.Add("want", foreign)
	req.URL.RawQuery = q.Encode()
}
