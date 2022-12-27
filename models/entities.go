package models

type Rate struct {
	Spot float64
}

type FXRate map[Pair]Rate

type Pair struct {
	Base    string
	Foreign string
}

type Amount struct {
	Unit     float64
	Currency Currency
}

type Currency struct {
	Code string
}
