package blackboxtests

import (
	"encoding/json"
	"github.com/mikejeuga/currency_converter/config"
	"github.com/mikejeuga/currency_converter/models"
	"io"
	"net/http"
	"net/url"
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

func (u *TestUser) GetRate(base, foreign string) (models.Rate, error) {
	rateURL, err := url.JoinPath(u.config.BaseURL, "rate")
	if err != nil {
		return models.Rate{}, err
	}

	req, err := http.NewRequest(http.MethodGet, rateURL, nil)
	if err != nil {
		return models.Rate{}, err
	}

	req.Header.Set("X-Api-Key", u.config.ApiKey)

	addQueryParams(req, base, foreign)

	res, err := u.client.Do(req)
	if err != nil {
		return models.Rate{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Rate{}, err
	}

	var rateRes models.Rate
	err = json.Unmarshal(data, &rateRes)
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

	req.Header.Set("X-Api-Key", u.config.ApiKey)

	addQueryParams(req, amount.Currency.Code, foreignCurrency)

	res, err := u.client.Do(req)
	data, err := io.ReadAll(res.Body)

	var returnedAmount models.Amount
	err = json.Unmarshal(data, &returnedAmount)
	if err != nil {
		return models.Amount{}, err
	}

	return returnedAmount, nil

}

func addQueryParams(req *http.Request, base, foreign string) {
	q := req.URL.Query()
	q.Add("have", base)
	q.Add("want", foreign)
	req.URL.RawQuery = q.Encode()
}
