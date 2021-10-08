package provider

import "monitoring-service/internal/app/model/entity"

// CryptocurrencyAPI define method for integration external APi
type CryptocurrencyAPI interface {
	GetRates(currencyID string) (*entity.Cryptocurrency, error)
}
