package blackboxtests

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mikejeuga/currency_converter/config"
	"github.com/mikejeuga/currency_converter/specifications"
	"log"
	"testing"
)

func TestAPI(t *testing.T) {
	var c config.Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	testUser := NewTestUser(c)
	spec := specifications.NewCurrencyConversionSpec(testUser)
	spec.CanConverterBaseToForeign(t)
}
