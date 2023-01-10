package client_test

import (
	"github.com/alecthomas/assert"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/currency_converter/blackboxtests/client"
	"github.com/mikejeuga/currency_converter/config"
	"github.com/mikejeuga/currency_converter/models"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	var cTest config.Config
	err := envconfig.Process("", &cTest)
	if err != nil {
		log.Fatal(err.Error())
	}

	testUser := client.NewTestUser(cTest)

	_, err = testUser.GetFXRate(models.CHF, models.USD)
	assert.NoError(t, err)
}
