//+go:build unit

package web_test

import (
	"github.com/alecthomas/assert"
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/specifications/mocks"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"github.com/mikejeuga/currency_converter/src/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	deps := CreateDeps()
	gateway := currency_conversion.NewGateway(deps.GatewayMock, currency_conversion.NewService())
	server := web.NewServer(gateway)

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
	} {
		t.Run(tc.description, func(t *testing.T) {
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
	GatewayMock *mocks.ConverterMock
}

func CreateDeps() Deps {
	return Deps{
		GatewayMock: &mocks.ConverterMock{},
	}
}
