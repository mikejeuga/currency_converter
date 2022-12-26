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
var _ currency_conversion.Converter = &ConverterMock{}

// ConverterMock is a mock implementation of currency_conversion.Converter.
//
//	func TestSomethingThatUsesConverter(t *testing.T) {
//
//		// make and configure a mocked currency_conversion.Converter
//		mockedConverter := &ConverterMock{
//			ConvertFunc: func(amount models.Amount, foreignCurrency string) (models.Amount, error) {
//				panic("mock out the Convert method")
//			},
//			GetRateFunc: func(base string, foreign string) (models.Rate, error) {
//				panic("mock out the GetRate method")
//			},
//		}
//
//		// use mockedConverter in code that requires currency_conversion.Converter
//		// and then make assertions.
//
//	}
type ConverterMock struct {
	// ConvertFunc mocks the Convert method.
	ConvertFunc func(amount models.Amount, foreignCurrency string) (models.Amount, error)

	// GetRateFunc mocks the GetRate method.
	GetRateFunc func(base string, foreign string) (models.Rate, error)

	// calls tracks calls to the methods.
	calls struct {
		// Convert holds details about calls to the Convert method.
		Convert []struct {
			// Amount is the amount argument value.
			Amount models.Amount
			// ForeignCurrency is the foreignCurrency argument value.
			ForeignCurrency string
		}
		// GetRate holds details about calls to the GetRate method.
		GetRate []struct {
			// Base is the base argument value.
			Base string
			// Foreign is the foreign argument value.
			Foreign string
		}
	}
	lockConvert sync.RWMutex
	lockGetRate sync.RWMutex
}

// Convert calls ConvertFunc.
func (mock *ConverterMock) Convert(amount models.Amount, foreignCurrency string) (models.Amount, error) {
	if mock.ConvertFunc == nil {
		panic("ConverterMock.ConvertFunc: method is nil but Converter.Convert was just called")
	}
	callInfo := struct {
		Amount          models.Amount
		ForeignCurrency string
	}{
		Amount:          amount,
		ForeignCurrency: foreignCurrency,
	}
	mock.lockConvert.Lock()
	mock.calls.Convert = append(mock.calls.Convert, callInfo)
	mock.lockConvert.Unlock()
	return mock.ConvertFunc(amount, foreignCurrency)
}

// ConvertCalls gets all the calls that were made to Convert.
// Check the length with:
//
//	len(mockedConverter.ConvertCalls())
func (mock *ConverterMock) ConvertCalls() []struct {
	Amount          models.Amount
	ForeignCurrency string
} {
	var calls []struct {
		Amount          models.Amount
		ForeignCurrency string
	}
	mock.lockConvert.RLock()
	calls = mock.calls.Convert
	mock.lockConvert.RUnlock()
	return calls
}

// GetRate calls GetRateFunc.
func (mock *ConverterMock) GetRate(base string, foreign string) (models.Rate, error) {
	if mock.GetRateFunc == nil {
		panic("ConverterMock.GetRateFunc: method is nil but Converter.GetRate was just called")
	}
	callInfo := struct {
		Base    string
		Foreign string
	}{
		Base:    base,
		Foreign: foreign,
	}
	mock.lockGetRate.Lock()
	mock.calls.GetRate = append(mock.calls.GetRate, callInfo)
	mock.lockGetRate.Unlock()
	return mock.GetRateFunc(base, foreign)
}

// GetRateCalls gets all the calls that were made to GetRate.
// Check the length with:
//
//	len(mockedConverter.GetRateCalls())
func (mock *ConverterMock) GetRateCalls() []struct {
	Base    string
	Foreign string
} {
	var calls []struct {
		Base    string
		Foreign string
	}
	mock.lockGetRate.RLock()
	calls = mock.calls.GetRate
	mock.lockGetRate.RUnlock()
	return calls
}
