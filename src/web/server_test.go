package web_test

import (
	"github.com/alecthomas/assert"
	"github.com/mikejeuga/currency_converter/src/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	server := web.NewServer()

	server.Handler.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}
