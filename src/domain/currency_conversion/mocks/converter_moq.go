// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/mikejeuga/currency_converter/models"
	"github.com/mikejeuga/currency_converter/src/domain/currency_conversion"
	"sync"
)

// Ensure, that ConverterMock does implement currency_conversion.Converter.
// If this is not the case, regenerate this file with moq.
var _ currency_conversion.Conversioner = &ConverterMock{}

// ConverterMock is a mock implementation of currency_conversion.Converter.
//
//	func TestSomethingThatUsesConverter(t *testing.T) {
//
//		// make and configure a mocked currency_conversion.Converter
//		mockedConverter := &ConverterMock{
//			GetFXRateFunc: func(base string, foreign string) (models.Rate, error) {
//				panic("mock out the GetFXRate method")
//			},
//		}
//
//		// use mockedConverter in code that requires currency_conversion.Converter
//		// and then make assertions.
//
//	}
type ConverterMock struct {
	// GetFXRateFunc mocks the GetFXRate method.
	GetFXRateFunc func(base string, foreign string) (models.Rate, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetFXRate holds details about calls to the GetFXRate method.
		GetFXRate []struct {
			// Base is the base argument value.
			Base string
			// Foreign is the foreign argument value.
			Foreign string
		}
	}
	lockGetFXRate sync.RWMutex
}

// GetFXRate calls GetFXRateFunc.
func (mock *ConverterMock) GetFXRate(base string, foreign string) (models.Rate, error) {
	if mock.GetFXRateFunc == nil {
		panic("ConverterMock.GetFXRateFunc: method is nil but Converter.GetFXRate was just called")
	}
	callInfo := struct {
		Base    string
		Foreign string
	}{
		Base:    base,
		Foreign: foreign,
	}
	mock.lockGetFXRate.Lock()
	mock.calls.GetFXRate = append(mock.calls.GetFXRate, callInfo)
	mock.lockGetFXRate.Unlock()
	return mock.GetFXRateFunc(base, foreign)
}

// GetFXRateCalls gets all the calls that were made to GetFXRate.
// Check the length with:
//
//	len(mockedConverter.GetFXRateCalls())
func (mock *ConverterMock) GetFXRateCalls() []struct {
	Base    string
	Foreign string
} {
	var calls []struct {
		Base    string
		Foreign string
	}
	mock.lockGetFXRate.RLock()
	calls = mock.calls.GetFXRate
	mock.lockGetFXRate.RUnlock()
	return calls
}
