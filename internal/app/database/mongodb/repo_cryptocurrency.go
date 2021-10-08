package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"monitoring-service/internal/app/model/entity"
	"time"
)

type cryptocurrencyRepository struct {
	collection *mongo.Collection
	ctxDur     time.Duration
}

// CreateCryptocurrency implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) CreateCryptocurrency(currency *entity.Cryptocurrency) error {
	id := primitive.NewObjectID()
	currency.ID = id
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxDur)
	defer cancel()
	_, err := c.collection.InsertOne(ctx, currency)
	if err != nil {
		return err
	}
	return nil
}

// GetCryptocurrencyByCurrencyID implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) GetCryptocurrencyByCurrencyID(currencyID string) (*entity.Cryptocurrency, error) {
	var cur *entity.Cryptocurrency
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxDur)
	defer cancel()
	filter := bson.M{"id": currencyID}
	if err := c.collection.FindOne(ctx, filter).Decode(&cur); err != nil {
		return nil, err
	}
	return cur, nil
}

// DeleteCryptocurrencyByCurrencyID implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) DeleteCryptocurrencyByCurrencyID(currencyID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxDur)
	defer cancel()
	filter := bson.M{"id": currencyID}
	_, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return err
}

// ListCryptocurrencies implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) ListCryptocurrencies(offset int, limit int) ([]*entity.Cryptocurrency, error) {
	currencies := make([]*entity.Cryptocurrency, 0, limit)
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxDur)
	defer cancel()
	skip := int64(offset)
	lim := int64(limit)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &lim,
	}
	cur, err := c.collection.Find(ctx, bson.M{}, &opts)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var currency *entity.Cryptocurrency
		if err := cur.Decode(&currency); err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}
	return currencies, nil
}

// ListCryptocurrenciesForUpdate implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) ListCryptocurrenciesForUpdate() ([]*entity.Cryptocurrency, error) {
	currencies := make([]*entity.Cryptocurrency, 0)
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxDur)
	defer cancel()
	filter := bson.M{"update_at": bson.M{"$lte": time.Now()}}
	cur, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var currency *entity.Cryptocurrency
		if err := cur.Decode(&currency); err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}
	return currencies, nil
}

// UpdateCryptocurrency implements database.CryptocurrencyRepository interface method
func (c *cryptocurrencyRepository) UpdateCryptocurrency(currency *entity.Cryptocurrency) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxDur)
	defer cancel()
	filter := bson.M{"_id": currency.ID}
	upd := bson.M{"rate_usd": currency.RateUsd, "updated": currency.Updated, "update_at": currency.UpdateAt}
	_, err := c.collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: upd}})
	if err != nil {
		return err
	}
	return nil
}

func newCryptocurrencyRepository(collection *mongo.Collection, ctxDur time.Duration) *cryptocurrencyRepository {
	return &cryptocurrencyRepository{collection: collection, ctxDur: ctxDur}
}
