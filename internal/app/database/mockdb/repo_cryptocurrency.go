package mockdb

import (
	"errors"
	"monitoring-service/internal/app/model/entity"
	"time"
)

type cryptocurrencyRepository struct {
	currencies []*entity.Cryptocurrency
	raiseError bool
}

// CreateCryptocurrency implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) CreateCryptocurrency(currency *entity.Cryptocurrency) error {
	if c.raiseError {
		return errors.New("failed to create")
	}
	for _, v := range c.currencies {
		if v.CurrencyId == currency.CurrencyId {
			return errors.New("already exists")
		}
	}
	c.currencies = append(c.currencies, currency)
	return nil
}

// GetCryptocurrencyByCurrencyID implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) GetCryptocurrencyByCurrencyID(currencyID string) (*entity.Cryptocurrency, error) {
	var result *entity.Cryptocurrency
	if c.raiseError {
		return nil, errors.New("failed to get currency")
	}
	for _, v := range c.currencies {
		if v.CurrencyId == currencyID {
			result = v
			break
		}
	}
	if result == nil {
		return nil, errors.New("currency not found")
	}
	return result, nil
}

// DeleteCryptocurrencyByCurrencyID implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) DeleteCryptocurrencyByCurrencyID(currencyID string) error {
	var result []*entity.Cryptocurrency
	if c.raiseError {
		return errors.New("failed to delete currency")
	}
	for _, v := range c.currencies {
		if v.CurrencyId != currencyID {
			result = append(result, v)
		}
	}
	c.currencies = result
	return nil
}

// ListCryptocurrencies implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) ListCryptocurrencies(offset, limit int) ([]*entity.Cryptocurrency, error) {
	var result []*entity.Cryptocurrency
	if c.raiseError {
		return nil, errors.New("failed to get list of currencies")
	}
	for i, v := range c.currencies {
		if i >= offset && i < limit {
			result = append(result, v)
		}
	}
	return result, nil
}

// ListCryptocurrenciesForUpdate implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) ListCryptocurrenciesForUpdate() ([]*entity.Cryptocurrency, error) {
	var result []*entity.Cryptocurrency
	if c.raiseError {
		return nil, errors.New("failed to get list of currencies")
	}
	for _, v := range c.currencies {
		if time.Now().After(v.UpdateAt) {
			result = append(result, v)
		}
	}
	return result, nil
}

// UpdateCryptocurrency implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) UpdateCryptocurrency(currency *entity.Cryptocurrency) error {
	if c.raiseError {
		return errors.New("failed to update currency")
	}
	for _, v := range c.currencies {
		if v.CurrencyId == currency.CurrencyId {
			v = currency
			return nil
		}
	}
	return errors.New("currency not found")
}

func newCryptocurrencyRepository(raiseError bool) *cryptocurrencyRepository {
	btc := entity.Cryptocurrency{
		CurrencyId:      "bitcoin",
		Symbol:          "BTC",
		CurrencySymbol:  "₿",
		Type:            "crypto",
		RateUsd:         "54405.3668813521278710",
		RefreshInterval: 180000000000,
		Updated:         time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
		UpdateAt:        time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
	}
	eth := entity.Cryptocurrency{
		CurrencyId:      "ethereum",
		Symbol:          "ETH",
		CurrencySymbol:  "E",
		Type:            "crypto",
		RateUsd:         "54405.3668813521278710",
		RefreshInterval: 180000000000,
		Updated:         time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
		UpdateAt:        time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
	}
	currencies := []*entity.Cryptocurrency{&btc, &eth}
	return &cryptocurrencyRepository{raiseError: raiseError, currencies: currencies}
}
