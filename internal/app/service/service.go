package service

import (
	"monitoring-service/internal/app/model/entity"
	"time"
)

// Service application service layer interface
type Service interface {
	Cryptocurrency() CryptocurrencyService
	Auth() AuthService
}

// CryptocurrencyService cryptocurrency service layer interface
type CryptocurrencyService interface {
	CreateCryptocurrency(currencyID string, refreshInterval time.Duration) (*entity.Cryptocurrency, error)
	GetCryptocurrencyByCurrencyID(currencyID string) (*entity.Cryptocurrency, error)
	DeleteCryptocurrencyByCurrencyID(currencyID string) (*entity.Cryptocurrency, error)
	ListCryptocurrencies(offset, limit int) ([]*entity.Cryptocurrency, error)
	UpdateCryptocurrencies() error
}

// AuthService JWT authentication service layer interface
type AuthService interface {
	Login(username, password string) (*entity.Tokens, error)
}
