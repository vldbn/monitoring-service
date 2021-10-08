package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCryptocurrencyService_CreateCryptocurrency(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		cur, err := svc.Cryptocurrency().CreateCryptocurrency("usd", time.Minute)
		assert.NoError(t, err)
		assert.NotNil(t, cur)
		assert.Equal(t, "usd", cur.CurrencyId)
	})
	t.Run("duplicate-currency", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		cur, err := svc.Cryptocurrency().CreateCryptocurrency("bitcoin", time.Minute)
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
	t.Run("db-error", func(t *testing.T) {
		svc := ConfigureServicesForTesting(true)
		cur, err := svc.Cryptocurrency().CreateCryptocurrency("bitcoin", time.Minute)
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}

func TestCryptocurrencyService_GetCryptocurrencyByCurrencyID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		cur, err := svc.Cryptocurrency().GetCryptocurrencyByCurrencyID("bitcoin")
		assert.NoError(t, err)
		assert.NotNil(t, cur)
		assert.Equal(t, "bitcoin", cur.CurrencyId)
	})
	t.Run("no-currency-in-db", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		cur, err := svc.Cryptocurrency().GetCryptocurrencyByCurrencyID("bitcoin1")
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
	t.Run("db-error", func(t *testing.T) {
		svc := ConfigureServicesForTesting(true)
		cur, err := svc.Cryptocurrency().GetCryptocurrencyByCurrencyID("bitcoin1")
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}

func TestCryptocurrencyService_DeleteCryptocurrencyByCurrencyID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		cur, err := svc.Cryptocurrency().DeleteCryptocurrencyByCurrencyID("bitcoin")
		assert.NoError(t, err)
		assert.NotNil(t, cur)
		assert.Equal(t, "bitcoin", cur.CurrencyId)
	})
	t.Run("no-currency-in-db", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		cur, err := svc.Cryptocurrency().DeleteCryptocurrencyByCurrencyID("bitcoin1")
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
	t.Run("db-error", func(t *testing.T) {
		svc := ConfigureServicesForTesting(true)
		cur, err := svc.Cryptocurrency().DeleteCryptocurrencyByCurrencyID("bitcoin1")
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}

func TestCryptocurrencyService_ListCryptocurrencies(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		cur, err := svc.Cryptocurrency().ListCryptocurrencies(0, 1)
		assert.NoError(t, err)
		assert.NotNil(t, cur)
	})
	t.Run("db-error", func(t *testing.T) {
		svc := ConfigureServicesForTesting(true)
		cur, err := svc.Cryptocurrency().ListCryptocurrencies(0, 1)
		assert.Error(t, err)
		assert.Nil(t, cur)
	})
}

func TestCryptocurrencyService_UpdateCryptocurrencies(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		err := svc.Cryptocurrency().UpdateCryptocurrencies()
		assert.NoError(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		svc := ConfigureServicesForTesting(true)
		err := svc.Cryptocurrency().UpdateCryptocurrencies()
		assert.Error(t, err)
	})
}
