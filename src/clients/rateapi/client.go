package rateapi

import (
	"encoding/json"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/web/auth"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	config Config
	client *http.Client
}

type Config struct {
	ApiURL string `envconfig:"API_URL"`
	//ApiKey string `envconfig:"X_API_KEY"`
}

func NewClient(config Config) *Client {
	c := &http.Client{
		Transport: auth.MyRoundTripper{Next: http.DefaultTransport},
		Timeout:   10 * time.Second,
	}
	return &Client{config: config, client: c}
}

func (c *Client) GetFXRate(base, foreign string) (models.Rate, error) {
	rateURL, err := url.JoinPath(c.config.ApiURL, "convertcurrency")
	if err != nil {
		return models.Rate{}, err
	}

	req, err := http.NewRequest(http.MethodGet, rateURL, nil)
	if err != nil {
		return models.Rate{}, err
	}

	addQueryParams(req, base, foreign)

	res, err := c.client.Do(req)
	if err != nil {
		return models.Rate{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Rate{}, err
	}

	var resRate RateResponse
	err = json.Unmarshal(data, &resRate)
	if err != nil {
		return models.Rate{}, err
	}

	return models.Rate{Spot: resRate.NewAmount / resRate.OldAmount}, nil
}

func addQueryParams(req *http.Request, base, foreign string) {
	q := req.URL.Query()
	q.Add("have", base)
	q.Add("want", foreign)
	q.Add("amount", "1")
	req.URL.RawQuery = q.Encode()
}
