package mockdb

import (
	"github.com/stretchr/testify/assert"
	"monitoring-service/internal/app/model/entity"
	"testing"
	"time"
)

func TestCryptocurrencyRepository_CreateCryptocurrency(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := NewMockDB(false)
		coin := entity.Cryptocurrency{
			CurrencyId:      "coin",
			Symbol:          "coin",
			CurrencySymbol:  "E",
			Type:            "crypto",
			RateUsd:         "54405.3668813521278710",
			RefreshInterval: 180000000000,
			Updated:         time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
			UpdateAt:        time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
		}
		err := db.CryptocurrencyRepo().CreateCryptocurrency(&coin)
		assert.NoError(t, err)
	})
	t.Run("failure-duplicate", func(t *testing.T) {
		db := NewMockDB(false)
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
		err := db.CryptocurrencyRepo().CreateCryptocurrency(&btc)
		assert.Error(t, err)
	})
}

func TestCryptocurrencyRepository_GetCryptocurrencyByCurrencyID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := NewMockDB(false)
		cur, err := db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID("bitcoin")
		assert.NoError(t, err)
		assert.NotNil(t, cur)
	})
	t.Run("failure-not-found", func(t *testing.T) {
		db := NewMockDB(false)
		cur, err := db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID("bitcoin22")
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}

func TestCryptocurrencyRepository_DeleteCryptocurrencyByCurrencyID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := NewMockDB(false)
		err := db.CryptocurrencyRepo().DeleteCryptocurrencyByCurrencyID("bitcoin")
		assert.NoError(t, err)
	})
	t.Run("failure-not-found", func(t *testing.T) {
		db := NewMockDB(false)
		err := db.CryptocurrencyRepo().DeleteCryptocurrencyByCurrencyID("bitcoin22")
		assert.NoError(t, err)
	})
}

func TestCryptocurrencyRepository_ListCryptocurrencies(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := NewMockDB(false)
		cur, err := db.CryptocurrencyRepo().ListCryptocurrencies(0, 3)
		assert.NoError(t, err)
		assert.NotNil(t, cur)
		assert.Equal(t, 2, len(cur))
	})
}

func TestCryptocurrencyRepository_ListCryptocurrenciesForUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := NewMockDB(false)
		cur, err := db.CryptocurrencyRepo().ListCryptocurrenciesForUpdate()
		assert.NoError(t, err)
		assert.NotNil(t, cur)
		assert.Equal(t, 2, len(cur))
	})
}

func TestCryptocurrencyRepository_UpdateCryptocurrency(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		newRate := "1000"
		db := NewMockDB(false)
		btc := entity.Cryptocurrency{
			CurrencyId:      "bitcoin",
			Symbol:          "BTC",
			CurrencySymbol:  "₿",
			Type:            "crypto",
			RateUsd:         newRate,
			RefreshInterval: 180000000000,
			Updated:         time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
			UpdateAt:        time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
		}
		err := db.CryptocurrencyRepo().UpdateCryptocurrency(&btc)
		assert.NoError(t, err)
		assert.Equal(t, newRate, btc.RateUsd)
	})
	t.Run("failure-not-found", func(t *testing.T) {
		newRate := "1000"
		db := NewMockDB(false)
		btc := entity.Cryptocurrency{
			CurrencyId:      "bitcoin2",
			Symbol:          "BTC",
			CurrencySymbol:  "₿",
			Type:            "crypto",
			RateUsd:         newRate,
			RefreshInterval: 180000000000,
			Updated:         time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
			UpdateAt:        time.Date(2021, 10, 7, 9, 59, 20, 123, time.UTC),
		}
		err := db.CryptocurrencyRepo().UpdateCryptocurrency(&btc)
		assert.Error(t, err)
	})
}
