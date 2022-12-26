package rateapi

type RateResponse struct {
	OldAmount   int     `json:"old_amount"`
	OldCurrency string  `json:"old_currency"`
	NewCurrency string  `json:"new_currency"`
	NewAmount   float64 `json:"new_amount"`
}
