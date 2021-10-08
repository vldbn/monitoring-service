package impl

import (
	"errors"
	"log"
	"monitoring-service/internal/app/database"
	"monitoring-service/internal/app/model/entity"
	"monitoring-service/internal/app/provider"
	"time"
)

type cryptocurrencyService struct {
	db        database.Database
	cryptoAPI provider.CryptocurrencyAPI
}

// CreateCryptocurrency implements service.CryptocurrencyService interface method
func (c *cryptocurrencyService) CreateCryptocurrency(
	currencyID string,
	refreshInterval time.Duration,
) (*entity.Cryptocurrency, error) {
	existCur, _ := c.db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID(currencyID)
	if existCur != nil {
		return nil, errors.New("currency already exists in database")
	}
	cur, err := c.cryptoAPI.GetRates(currencyID)
	if err != nil {
		return nil, err
	}
	createdAt := time.Now()
	updateAt := createdAt.Add(refreshInterval * time.Minute)
	cur.Updated = createdAt
	cur.UpdateAt = updateAt
	cur.RefreshInterval = refreshInterval * time.Minute
	if err := c.db.CryptocurrencyRepo().CreateCryptocurrency(cur); err != nil {
		return nil, err
	}
	return cur, nil
}

// GetCryptocurrencyByCurrencyID implements service.CryptocurrencyService interface method
func (c *cryptocurrencyService) GetCryptocurrencyByCurrencyID(currencyID string) (*entity.Cryptocurrency, error) {
	cur, err := c.db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID(currencyID)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

// DeleteCryptocurrencyByCurrencyID implements service.CryptocurrencyService interface method
func (c *cryptocurrencyService) DeleteCryptocurrencyByCurrencyID(currencyID string) (*entity.Cryptocurrency, error) {
	cur, err := c.db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID(currencyID)
	if err != nil {
		return nil, err
	}
	if err := c.db.CryptocurrencyRepo().DeleteCryptocurrencyByCurrencyID(cur.CurrencyId); err != nil {
		return nil, err
	}
	return cur, nil
}

// ListCryptocurrencies implements service.CryptocurrencyService interface method
func (c *cryptocurrencyService) ListCryptocurrencies(offset, limit int) ([]*entity.Cryptocurrency, error) {
	cur, err := c.db.CryptocurrencyRepo().ListCryptocurrencies(offset, limit)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

// UpdateCryptocurrencies implements service.CryptocurrencyService interface method
func (c *cryptocurrencyService) UpdateCryptocurrencies() error {
	currencies, err := c.db.CryptocurrencyRepo().ListCryptocurrenciesForUpdate()
	if err != nil {
		return err
	}
	log.Printf("Schedule task: %d curencies for update", len(currencies))
	for _, currency := range currencies {
		updCurrency, err := c.cryptoAPI.GetRates(currency.CurrencyId)
		if err != nil {
			return err
		}
		now := time.Now()
		currency.RateUsd = updCurrency.RateUsd
		currency.Updated = now
		currency.UpdateAt = now.Add(currency.RefreshInterval * time.Nanosecond)
		if err := c.db.CryptocurrencyRepo().UpdateCryptocurrency(currency); err != nil {
			return err
		}
		log.Printf("Schedule task: %s was updated. New rate - %s", currency.CurrencyId, currency.RateUsd)
	}
	return nil
}

func newCryptocurrencyService(db database.Database, cryptoAPI provider.CryptocurrencyAPI) *cryptocurrencyService {
	return &cryptocurrencyService{db: db, cryptoAPI: cryptoAPI}
}
