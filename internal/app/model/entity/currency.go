package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Cryptocurrency entity struct
type Cryptocurrency struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CurrencyId      string             `json:"id" bson:"id"`
	Symbol          string             `json:"symbol" bson:"symbol"`
	CurrencySymbol  string             `json:"currencySymbol" bson:"currency_symbol"`
	Type            string             `json:"type" bson:"type"`
	RateUsd         string             `json:"rateUsd" bson:"rate_usd"`
	RefreshInterval time.Duration      `json:"refresh_interval,omitempty" bson:"refresh_interval"`
	Updated         time.Time          `json:"updated,omitempty" bson:"updated"`
	UpdateAt        time.Time          `json:"update_at,omitempty" bson:"update_at"`
}
