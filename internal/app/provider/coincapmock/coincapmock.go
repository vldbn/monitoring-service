package coincapmock

import (
	"fmt"
	"math/rand"
	"monitoring-service/internal/app/model/entity"
)

// MockCoinCapAPI implements provider.CryptocurrencyAPI interface
type MockCoinCapAPI struct {
}

// GetRates implements provider.CryptocurrencyAPI interface method
func (m *MockCoinCapAPI) GetRates(currencyID string) (*entity.Cryptocurrency, error) {
	currency := entity.Cryptocurrency{
		CurrencyId:     currencyID,
		Symbol:         "CUR",
		CurrencySymbol: "C",
		Type:           "crypto",
		RateUsd:        fmt.Sprintf("%f", rand.ExpFloat64()),
	}
	return &currency, nil
}

// NewMockCoinCapAPI constructor
func NewMockCoinCapAPI() *MockCoinCapAPI {
	return &MockCoinCapAPI{}
}
