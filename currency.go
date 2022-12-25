package currency_converter

import "time"

type Rate float64

type Currency struct {
	Code         string
	OfficialName string
	Rate         Rate
	TimeStamp    time.Time
}

type Amount struct {
	Unit     float64
	Currency Currency
}

type Pair struct {
	Base      Currency
	Foreign   Currency
	TimeStamp time.Time
}

type FXRate map[Pair]Rate

const (
	GBP = "GBP"
	USD = "USD"
	EUR = "EUR"
	CHF = "CHF"
	JPY = "JPY"
	AUD = "AUD"
	SEK = "SEK"
	CNY = "CNY"
)
