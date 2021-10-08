package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"monitoring-service/internal/app/core"
	"time"
)

// SetMongoDBClientForTesting set mongo client for testing
func SetMongoDBClientForTesting() (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	config := core.NewConfig()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseURL()))
	return client, ctx, cancel
}
