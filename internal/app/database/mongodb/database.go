package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"monitoring-service/internal/app/database"
	"time"
)

// MongoDB implements database.Database interface
type MongoDB struct {
	client             *mongo.Client
	database           string
	ctxDur             time.Duration
	cryptocurrencyRepo *cryptocurrencyRepository
}

// CryptocurrencyRepo implements database.Database interface method
func (m *MongoDB) CryptocurrencyRepo() database.CryptocurrencyRepository {
	if m.cryptocurrencyRepo != nil {
		return m.cryptocurrencyRepo
	}
	collection := m.client.Database(m.database).Collection("cryptocurrencies")
	m.cryptocurrencyRepo = newCryptocurrencyRepository(collection, m.ctxDur)
	return m.cryptocurrencyRepo
}

// NewMongoDB constructor
func NewMongoDB(client *mongo.Client, database string) database.Database {
	return &MongoDB{
		client:   client,
		database: database,
		ctxDur:   time.Second * 3,
	}
}
