package mockdb

import (
	"monitoring-service/internal/app/database"
)

// MockDB implements database.Database interface
type MockDB struct {
	cryptocurrencyRepo *cryptocurrencyRepository
	raiseError         bool
}

// CryptocurrencyRepo  implements database.Database interface method
func (m *MockDB) CryptocurrencyRepo() database.CryptocurrencyRepository {
	if m.cryptocurrencyRepo != nil {
		return m.cryptocurrencyRepo
	}
	m.cryptocurrencyRepo = newCryptocurrencyRepository(m.raiseError)
	return m.cryptocurrencyRepo
}

// NewMockDB constructor
func NewMockDB(raiseError bool) *MockDB {
	return &MockDB{raiseError: raiseError}
}
