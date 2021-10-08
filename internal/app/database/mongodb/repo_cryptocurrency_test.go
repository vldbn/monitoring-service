package mongodb

import (
	"github.com/stretchr/testify/assert"
	"monitoring-service/internal/app/model/entity"
	"testing"
	"time"
)

const testingDB = "monitoring-testing"

func TestCryptocurrencyRepository_CreateCryptocurrency(t *testing.T) {
	client, ctx, cancel := SetMongoDBClientForTesting()
	defer func() {
		_ = client.Database(testingDB).Drop(ctx)
		cancel()
	}()
	t.Run("success", func(t *testing.T) {
		db := NewMongoDB(client, testingDB)
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
		assert.NoError(t, err)
	})
}

func TestCryptocurrencyRepository_GetCryptocurrencyByCurrencyID(t *testing.T) {
	client, ctx, cancel := SetMongoDBClientForTesting()
	defer func() {
		_ = client.Database(testingDB).Drop(ctx)
		cancel()
	}()
	t.Run("success", func(t *testing.T) {
		db := NewMongoDB(client, testingDB)
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
		_ = db.CryptocurrencyRepo().CreateCryptocurrency(&btc)
		cur, err := db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID("bitcoin")
		assert.NoError(t, err)
		assert.NotNil(t, cur)
	})
	t.Run("failure-not-found", func(t *testing.T) {
		db := NewMongoDB(client, testingDB)
		cur, err := db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID("bitcoin22")
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}

func TestCryptocurrencyRepository_DeleteCryptocurrencyByCurrencyID(t *testing.T) {
	client, ctx, cancel := SetMongoDBClientForTesting()
	defer func() {
		_ = client.Database(testingDB).Drop(ctx)
		cancel()
	}()
	t.Run("success", func(t *testing.T) {
		db := NewMongoDB(client, testingDB)
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
		_ = db.CryptocurrencyRepo().CreateCryptocurrency(&btc)
		err := db.CryptocurrencyRepo().DeleteCryptocurrencyByCurrencyID("bitcoin")
		assert.NoError(t, err)
		cur, err := db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID("bitcoin")
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}

func TestCryptocurrencyRepository_ListCryptocurrencies(t *testing.T) {
	client, ctx, cancel := SetMongoDBClientForTesting()
	defer func() {
		_ = client.Database(testingDB).Drop(ctx)
		cancel()
	}()
	t.Run("success", func(t *testing.T) {
		db := NewMongoDB(client, testingDB)
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
		_ = db.CryptocurrencyRepo().CreateCryptocurrency(&btc)
		currencies, err := db.CryptocurrencyRepo().ListCryptocurrencies(0, 5)
		assert.NoError(t, err)
		assert.NotNil(t, currencies)
		assert.Equal(t, 1, len(currencies))
	})
}

func TestCryptocurrencyRepository_ListCryptocurrenciesForUpdate(t *testing.T) {
	client, ctx, cancel := SetMongoDBClientForTesting()
	defer func() {
		_ = client.Database(testingDB).Drop(ctx)
		cancel()
	}()
	t.Run("success", func(t *testing.T) {
		db := NewMongoDB(client, testingDB)
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
		_ = db.CryptocurrencyRepo().CreateCryptocurrency(&btc)
		currencies, err := db.CryptocurrencyRepo().ListCryptocurrenciesForUpdate()
		assert.NoError(t, err)
		assert.NotNil(t, currencies)
		assert.Equal(t, 1, len(currencies))
	})
}

func TestCryptocurrencyRepository_UpdateCryptocurrency(t *testing.T) {
	client, ctx, cancel := SetMongoDBClientForTesting()
	defer func() {
		_ = client.Database(testingDB).Drop(ctx)
		cancel()
	}()
	t.Run("success", func(t *testing.T) {
		db := NewMongoDB(client, testingDB)
		newRate := "10000"
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
		_ = db.CryptocurrencyRepo().CreateCryptocurrency(&btc)
		btc.RateUsd = newRate
		err := db.CryptocurrencyRepo().UpdateCryptocurrency(&btc)
		assert.NoError(t, err)
		cur, _ := db.CryptocurrencyRepo().GetCryptocurrencyByCurrencyID("bitcoin")
		assert.Equal(t, newRate, cur.RateUsd)
	})
}
