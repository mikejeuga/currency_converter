package blackboxtests

import (
	"github.com/mikejeuga/currency_converter/specifications"
	"testing"
)

func TestAPI(t *testing.T) {
	testUser := NewTestUser("http://localhost:8087")
	spec := specifications.NewCurrencyConversionSpec(testUser)
	spec.CanConverterBaseToForeign(t)
}
