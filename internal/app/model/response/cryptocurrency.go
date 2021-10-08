package response

import "monitoring-service/internal/app/model/entity"

// CryptocurrencyRes cryptocurrency HTTP response struct
type CryptocurrencyRes struct {
	DefaultRes
	Cryptocurrency *entity.Cryptocurrency `json:"cryptocurrency,omitempty"`
}

// CryptocurrenciesRes cryptocurrencies HTTP response struct
type CryptocurrenciesRes struct {
	DefaultRes
	Cryptocurrencies []*entity.Cryptocurrency `json:"cryptocurrencies,omitempty"`
}
