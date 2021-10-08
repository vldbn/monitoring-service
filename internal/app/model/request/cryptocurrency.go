package request

import "time"

// CreateCryptocurrencyReq cryptocurrency create request struct
type CreateCryptocurrencyReq struct {
	CurrencyID      string        `json:"id"`
	RefreshInterval time.Duration `json:"refresh_interval"`
}
