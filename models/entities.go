package models

import "time"

type Rate struct {
	Spot    float64
	Forward float64
}

type FXRate map[Pair]Rate

type Pair struct {
	Base      string
	Foreign   string
	TimeStamp time.Time
}

type Amount struct {
	Unit     float64
	Currency Currency
}

type Currency struct {
	Code string
}
