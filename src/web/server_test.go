//+go:build unit

package web_test

import (
	"github.com/alecthomas/assert"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/currency_converter/config"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"github.com/mikejeuga/currency_converter/src/web"
	mocks2 "github.com/mikejeuga/currency_converter/src/web/mocks"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	deps := CreateDeps()
	var testConf config.Config
	err := envconfig.Process("", &testConf)
	if err != nil {
		log.Fatal(err.Error())
	}

	gateway := currency_conversion.NewGateway(deps.GatewayMock, currency_conversion.NewService())
	server := web.NewServer(testConf, gateway)

	for _, tc := range []struct {
		description    string
		res            *httptest.ResponseRecorder
		req            *http.Request
		ExpectedStatus int
	}{
		{
			description:    "Route to Home endpoint '/'",
			res:            httptest.NewRecorder(),
			req:            httptest.NewRequest(http.MethodGet, "/", nil),
			ExpectedStatus: http.StatusOK,
		},
		{
			description:    "Route to Rate endpoint '/rate'",
			res:            httptest.NewRecorder(),
			req:            httptest.NewRequest(http.MethodGet, "/rate", nil),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			description:    "Route to Rate endpoint '/rate?have=GBP&want=USD'",
			res:            httptest.NewRecorder(),
			req:            httptest.NewRequest(http.MethodGet, "/rate?have=GBP&want=USD", nil),
			ExpectedStatus: http.StatusOK,
		},
		{
			description:    "Route to Rate endpoint '/converted-amount'",
			res:            httptest.NewRecorder(),
			req:            httptest.NewRequest(http.MethodGet, "/converted-amount", nil),
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			description:    "Route to Rate endpoint '/converted-amount?amount=1000&have=GBP&want=USD'",
			res:            httptest.NewRecorder(),
			req:            httptest.NewRequest(http.MethodGet, "/converted-amount?amount=1000&have=GBP&want=USD", nil),
			ExpectedStatus: http.StatusOK,
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			tc.req.Header.Set("X-API-KEY", testConf.ApiKey)
			givenGetRateWasCalled(deps)
			server.Handler.ServeHTTP(tc.res, tc.req)

			assert.Equal(t, tc.ExpectedStatus, tc.res.Code)

		})
	}
}

func givenGetRateWasCalled(deps Deps) {
	deps.GatewayMock.GetFXRateFunc = func(base string, foreign string) (models.Rate, error) {
		return models.Rate{}, nil
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
