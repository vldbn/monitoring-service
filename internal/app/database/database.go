package database

import (
	"monitoring-service/internal/app/model/entity"
)

// Database persistence layer interface
type Database interface {
	CryptocurrencyRepo() CryptocurrencyRepository
}

// CryptocurrencyRepository database repository layer interface
type CryptocurrencyRepository interface {
	CreateCryptocurrency(currency *entity.Cryptocurrency) error
	GetCryptocurrencyByCurrencyID(currencyID string) (*entity.Cryptocurrency, error)
	DeleteCryptocurrencyByCurrencyID(currencyID string) error
	ListCryptocurrencies(offset, limit int) ([]*entity.Cryptocurrency, error)
	ListCryptocurrenciesForUpdate() ([]*entity.Cryptocurrency, error)
	UpdateCryptocurrency(currency *entity.Cryptocurrency) error
}
